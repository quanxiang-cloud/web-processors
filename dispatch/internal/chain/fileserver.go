package chain

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"github.com/quanxiang-cloud/cabin/logger"
	"github.com/quanxiang-cloud/cabin/tailormade/header"
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
	}
}

func (f *fileserver) Name() string {
	return "fileserver"
}

func (f *fileserver) Do(ctx context.Context, params *Parameter) error {
	var (
		stderr      bytes.Buffer
		commandPath = fmt.Sprintf("%s/%s", f.conf.CommandDir, f.Name())
	)

	uploadpath := fmt.Sprintf("%s/%s/%s", f.conf.UploadPrefix, "hash", params.File.Filename)
	cmd := exec.Command(
		commandPath,
		"-filePath", params.CssFilePath,
		"-uploadPath", uploadpath,
	)

	cmd.Stderr = &stderr
	params.StorePath = uploadpath

	if _, err := cmd.Output(); err != nil {
		logger.Logger.Error("Execute Fileserver", "err", err.Error(), header.GetRequestIDKV(ctx).Fuzzy())

		return err
	}

	if stderr.Len() > 0 {
		logger.Logger.Errorw("Execute Fileserver", header.GetRequestIDKV(ctx).Fuzzy()...)
	}

	return f.next.Do(ctx, params)
}
