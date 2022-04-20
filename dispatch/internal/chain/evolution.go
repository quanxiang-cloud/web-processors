package chain

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"strings"

	error2 "github.com/quanxiang-cloud/cabin/error"
	"github.com/quanxiang-cloud/cabin/logger"
	"github.com/quanxiang-cloud/cabin/tailormade/header"
	"github.com/quanxiang-cloud/web-processors/dispatch/pkg/code"
	"github.com/quanxiang-cloud/web-processors/dispatch/pkg/config"
)

type evolution struct {
	conf *config.Config
	next Command
}

// NewEvolution returns a new evolution command.
func NewEvolution(conf *config.Config) *evolution {
	return &evolution{
		next: NewFileserver(conf),
		conf: conf,
	}
}

func (e *evolution) Name() string {
	return genCommandPath(e.conf.CommandDir, EvolutionCommandName)
}

func (e *evolution) Do(ctx context.Context, params *Parameter) error {
	saveUploadPath, err := e.saveUploadedFile(params.File)
	if err != nil {
		logger.Logger.WithName("Save File").Errorw(err.Error(), header.GetRequestIDKV(ctx).Fuzzy()...)
		return error2.New(code.ErrSaveFile)
	}

	defer os.Remove(saveUploadPath)

	var (
		stderr      bytes.Buffer
		commandPath = e.Name()
	)

	cmd := &exec.Cmd{
		Path: commandPath,
		Env: []string{
			genEnv(INPUTFILE, saveUploadPath),
		},

		Stderr: &stderr,
	}

	stdout, err := cmd.Output()
	if err != nil {
		logger.Logger.WithName("Execute evolution").Errorw(err.Error(), header.GetRequestIDKV(ctx).Fuzzy()...)
		return error2.New(code.ErrExecute)
	}

	if stderr.Len() > 0 {
		logger.Logger.WithName("Execute evolution").Errorw(stderr.String(), header.GetRequestIDKV(ctx).Fuzzy()...)
		return error2.New(code.ErrExecute)
	}

	arr := strings.Split(string(stdout), ",")
	if len(arr) != 2 {
		logger.Logger.WithName("Execute evolution").Errorw("Execute evolution error", header.GetRequestIDKV(ctx).Fuzzy()...)
		return error2.New(code.ErrExecute)
	}

	params.UploadFilePath = saveUploadPath
	params.CssFilePath = arr[0]
	params.CssFileHash = arr[1]

	return e.next.Do(ctx, params)
}

func (e *evolution) saveUploadedFile(file *multipart.FileHeader) (string, error) {
	err := os.MkdirAll(e.conf.UploadDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	dst := fmt.Sprintf("%s/%s", e.conf.UploadDir, file.Filename)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return dst, err
}
