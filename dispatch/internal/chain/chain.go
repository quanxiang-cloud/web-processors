package chain

import (
	"context"
	"fmt"
	"mime/multipart"
)

// Command is the interface that wraps the basic 'Do' method.
type Command interface {
	Do(ctx context.Context, args *Parameter) error
}

// Parameter parameter.
type Parameter struct {
	AppID string                `form:"appID"`
	Token string                `form:"token" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`

	Universal
}

// Universal is the universal parameter.
type Universal struct {
	UploadFilePath string
	CssFilePath    string
	CssFileHash    string
	StorePath      string
}

// command name.
const (
	EvolutionCommandName  = "evolution"
	FileserverCommandName = "fileserver"
	PersonaCommandName    = "persona"
)

func genCommandPath(dir, name string) string {
	return fmt.Sprintf("%s/%s", dir, name)
}

// env name.
const (
	INPUTFILE = "INPUT_FILE"
)

func genEnv(key, value string) string {
	return fmt.Sprintf("%s=%s", key, value)
}
