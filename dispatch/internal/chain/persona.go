package chain

import (
	"context"
	"os/exec"
	"strings"

	error2 "github.com/quanxiang-cloud/cabin/error"
	"github.com/quanxiang-cloud/cabin/logger"
	"github.com/quanxiang-cloud/cabin/tailormade/header"
	"github.com/quanxiang-cloud/web-processors/dispatch/pkg/code"
	"github.com/quanxiang-cloud/web-processors/dispatch/pkg/config"
)

type persona struct {
	conf *config.Config
}

// NewPersona returns a new persona command.
func NewPersona(conf *config.Config) Command {
	return &persona{
		conf: conf,
	}
}

func (p *persona) Name() string {
	return genCommandPath(p.conf.CommandDir, PersonaCommandName)
}

func (p *persona) Do(ctx context.Context, params *Parameter) error {
	// generate persona key
	key := p.genPersonaKey(ctx, params)
	logger.Logger.Info("Persona key: ", key)

	cmd := &exec.Cmd{
		Path: p.Name(),
		Args: []string{
			p.Name(),
			"-endpoint", p.conf.PersonaEndpoint,
			"-version", p.conf.PersonaVersion,
			"-key", key,
			"-value", params.StorePath,
		},
	}

	if _, err := execute(cmd); err != nil {
		logger.Logger.WithName("Execute Persona").Errorw(err.Error(), header.GetRequestIDKV(ctx).Fuzzy()...)
		return error2.New(code.ErrExecute)
	}

	return nil
}

func (p *persona) genPersonaKey(ctx context.Context, params *Parameter) string {
	// _, tenantIDValue := header.GetTenantID(ctx).Wreck()
	key := strings.Join(p.removeEmptyStr([]string{
		// tenantIDValue,
		params.AppID,
		p.conf.PersonaKeySuffix,
	}), ":")

	return key
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
