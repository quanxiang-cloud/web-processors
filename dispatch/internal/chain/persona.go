package chain

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/quanxiang-cloud/cabin/logger"
	"github.com/quanxiang-cloud/cabin/tailormade/header"
	"github.com/quanxiang-cloud/web-processors/dispatch/pkg/config"
)

type persona struct {
	conf *config.Config
}

// NewPersona returns a new persona command.
func NewPersona(conf *config.Config) *persona {
	return &persona{}
}

func (p *persona) Name() string {
	return "persona"
}

func (p *persona) Do(ctx context.Context, params *Parameter) error {
	var (
		stderr      bytes.Buffer
		commandPath = fmt.Sprintf("%s/%s", p.conf.CommandDir, p.Name())
	)

	_, tenantIDValue := header.GetTenantID(ctx).Wreck()
	param := []string{params.AppID, tenantIDValue, "style_guide_css:draft"}
	key := strings.Join(p.removeEmptyStr(param), ":")

	cmd := exec.Command(
		commandPath,
		"-url", p.conf.PersonaURL,
		"-key", key,
		"-value", params.StorePath,
	)

	cmd.Stderr = &stderr

	if _, err := cmd.Output(); err != nil {
		logger.Logger.Error("Execute Persona", "err", err.Error(), header.GetRequestIDKV(ctx).Fuzzy())

		return err
	}

	if stderr.Len() > 0 {
		logger.Logger.Errorw("Execute Fileserver", header.GetRequestIDKV(ctx).Fuzzy()...)
	}

	return nil
}

func (p *persona) removeEmptyStr(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "" {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}

	return s
}
