package chain

import (
	"context"

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

func (p *persona) Do(ctx context.Context, args *Arg) error {
	return nil
}
