package chain

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"

	"github.com/quanxiang-cloud/cabin/logger"
	"github.com/quanxiang-cloud/cabin/tailormade/header"
	"github.com/quanxiang-cloud/web-processors/dispatch/pkg/config"
)

type scss struct {
	conf *config.Config
	next Command
}

// NewSCSS returns a new scss command.
func NewSCSS(conf *config.Config) *scss {
	return &scss{
		next: NewCSS(conf),
	}
}

func (s *scss) Name() string {
	return "scss"
}

func (s *scss) Do(ctx context.Context, args *Arg) error {
	err := os.MkdirAll(s.conf.UploadDir, 0o655)
	if err != nil {
		logger.Logger.WithName("Create Dir").Errorw(err.Error(), header.GetRequestIDKV(ctx).Fuzzy()...)
		return err
	}

	saveUploadPath := fmt.Sprintf("%s/%s", s.conf.UploadDir, args.File.Filename)
	err = s.saveUploadedFile(args.File, saveUploadPath)
	if err != nil {
		logger.Logger.WithName("Save File").Errorw(err.Error(), header.GetRequestIDKV(ctx).Fuzzy()...)
		return err
	}

	args.UploadFilePath = saveUploadPath

	var (
		stderr      bytes.Buffer
		commandPath = fmt.Sprintf("%s/%s", s.conf.CommandDir, s.Name())
	)

	cmd := exec.Command(
		commandPath,
		"-f", args.UploadFilePath,
	)
	cmd.Stderr = &stderr
	if _, err := cmd.Output(); err != nil {
		logger.Logger.Error("Execute Scss", "err", err.Error(), header.GetRequestIDKV(ctx).Fuzzy())
		return err
	}

	return s.next.Do(ctx, args)
}

func (s *scss) saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
