package chain

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	error2 "github.com/quanxiang-cloud/cabin/error"
	"github.com/quanxiang-cloud/cabin/logger"
	"github.com/quanxiang-cloud/cabin/tailormade/header"
	"github.com/quanxiang-cloud/web-processors/dispatch/pkg/code"
	"github.com/quanxiang-cloud/web-processors/dispatch/pkg/config"
)

type fileserver struct {
	conf *config.Config
	next Command
}

// NewFileserver returns a new fileserver command.
func NewFileserver(conf *config.Config) *fileserver {
	return &fileserver{
		next: NewPersona(conf),
		conf: conf,
	}
}

func (f *fileserver) Name() string {
	return genCommandPath(f.conf.CommandDir, FileserverCommandName)
}

func (f *fileserver) Do(ctx context.Context, params *Parameter) error {
	var (
		stderr      bytes.Buffer
		commandPath = f.Name()
		storePath   = f.genStorePath(params)
	)

	cmd := exec.Cmd{
		Path: commandPath,
		Args: []string{
			commandPath,
			"-filePath", params.CssFilePath,
			"-uploadPath", storePath,
		},
		Stderr: &stderr,
	}

	if _, err := cmd.Output(); err != nil {
		logger.Logger.WithName("Execute Fileserver").Errorw(err.Error(), header.GetRequestIDKV(ctx).Fuzzy()...)
		return error2.New(code.ErrExecute)
	}

	if stderr.Len() > 0 {
		logger.Logger.WithName("Execute Fileserver").Errorw(stderr.String(), header.GetRequestIDKV(ctx).Fuzzy()...)
		return error2.New(code.ErrExecute)
	}

	params.StorePath = storePath

	return f.next.Do(ctx, params)
}

func (f *fileserver) genStorePath(params *Parameter) string {
	return fmt.Sprintf(
		"%s/%s/%s",
		f.conf.UploadPrefix,
		params.CssFileHash,
		params.File.Filename,
	)
}
