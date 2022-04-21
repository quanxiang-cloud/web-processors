package chain

import (
	"context"
	"fmt"
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
	saveUploadPath, err := e.saveUploadedFile(params.Element)
	if err != nil {
		logger.Logger.WithName("Save File").Errorw(err.Error(), header.GetRequestIDKV(ctx).Fuzzy()...)
		return error2.New(code.ErrSaveFile)
	}
	defer os.Remove(saveUploadPath)

	cmd := &exec.Cmd{
		Path: e.Name(),
		Env: []string{
			genEnv(INPUT_FILE, saveUploadPath),
			genEnv(PERSONA_HOSTNAME, e.conf.PersonaEndpoint),
		},
	}

	stdout, err := execute(cmd)
	if err != nil {
		logger.Logger.WithName("Execute Evolution").Errorw(err.Error(), header.GetRequestIDKV(ctx).Fuzzy()...)
		return error2.New(code.ErrExecute)
	}

	arr := strings.Split(stdout, ",")
	if len(arr) != 2 {
		logger.Logger.WithName("Execute Evolution").Errorw("invalid stdout", header.GetRequestIDKV(ctx).Fuzzy()...)
		return error2.New(code.ErrExecute)
	}

	params.CssFileHash = arr[0]
	params.CssFilePath = arr[1]
	params.UploadFilePath = saveUploadPath

	return e.next.Do(ctx, params)
}

func (e *evolution) saveUploadedFile(elem string) (string, error) {
	err := os.MkdirAll(e.conf.UploadDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	dst := fmt.Sprintf("%s/%s", e.conf.UploadDir, uploadFileName)

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = out.WriteString(elem)
	return dst, err
}
