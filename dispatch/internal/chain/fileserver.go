package chain

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
func NewFileserver(conf *config.Config) Command {
	return &fileserver{
		next: NewPersona(conf),
		conf: conf,
	}
}

func (f *fileserver) Name() string {
	return genCommandPath(f.conf.CommandDir, FileserverCommandName)
}

func (f *fileserver) Do(ctx context.Context, params *Parameter) error {
	logger.Logger.Infof("template string %#v", params)

	commandPath := f.Name()
	storePath := f.genStorePath(params)

	logger.Logger.WithName("xxxx").Info("AAAAAAA", commandPath, storePath)

	defer os.RemoveAll(params.CSSFilePath)

	cmd := &exec.Cmd{
		Path: commandPath,
		Args: []string{
			commandPath,
			"-filePath", params.CSSFilePath,
			"-storePath", storePath,
		},
	}
	ss, err := execute(cmd)
	if err != nil {
		logger.Logger.WithName("Execute Fileserver").Errorw(err.Error(), header.GetRequestIDKV(ctx).Fuzzy()...)
		return error2.New(code.ErrExecute)
	}
	logger.Logger.WithName("Execute Fileserver").Info("fileserver result", ss)

	params.StorePath = storePath

	logger.Logger.WithName("xxxx").Info("zzz", params.StorePath)

	return f.next.Do(ctx, params)
}

func (f *fileserver) genStorePath(params *Parameter) string {
	logger.Logger.WithName("xxxx").Info("zzz", params.CSSFilePath)
	_, filename := filepath.Split(params.CSSFilePath)
	logger.Logger.WithName("xxxx").Info("css file name", filename)

	s := fmt.Sprintf(
		"%s/%s/%s",
		f.conf.StorePathPrefix,
		params.CSSFileHash,
		filename,
	)

	logger.Logger.WithName("xxxx").Info("store path", s)
	return fmt.Sprintf(
		"%s/%s/%s",
		f.conf.StorePathPrefix,
		params.CSSFileHash,
		filename,
	)
}
