package chain

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/quanxiang-cloud/cabin/logger"
	"github.com/quanxiang-cloud/cabin/tailormade/header"
	"github.com/quanxiang-cloud/web-processors/dispatch/pkg/config"
)

type css struct {
	conf *config.Config
	next Command
}

// NewCSS css.
func NewCSS(conf *config.Config) *css {
	return &css{
		conf: conf,
		next: NewFileserver(conf),
	}
}

func (c *css) Name() string {
	return "css"
}

func (c *css) Do(ctx context.Context, params *Parameter) error {
	var (
		stderr      bytes.Buffer
		commandPath = fmt.Sprintf("%s/%s", c.conf.CommandDir, c.Name())
	)

	cmd := exec.Command(commandPath)
	cmd.Stderr = &stderr

	if _, err := cmd.Output(); err != nil {
		logger.Logger.Error("execute css", "err", err.Error(), header.GetRequestIDKV(ctx).Fuzzy())
		return err
	}

	if stderr.Len() > 0 {
		logger.Logger.Error("execute css", "err", stderr.String(), header.GetRequestIDKV(ctx).Fuzzy())
		return nil
	}

	params.CssFilePath = os.Getenv("CSS_FILE_PATH")

	defer os.RemoveAll(params.CssFilePath)

	return c.next.Do(ctx, params)
}
