package chain

import (
	"context"
	"mime/multipart"
)

// Command is the interface that wraps the basic 'Do' method.
type Command interface {
	Do(ctx context.Context, args *Arg) error
}

type Arg struct {
	Token string                `form:"token" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`

	Universal
}

type Universal struct {
	// XXX: xxxxxxxx
	UploadPrefix   string
	UploadFilePath string
	CssFilePath    string
}
