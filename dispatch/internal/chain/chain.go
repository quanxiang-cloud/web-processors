package chain

import (
	"context"
	"mime/multipart"
)

// Command is the interface that wraps the basic 'Do' method.
type Command interface {
	Do(ctx context.Context, args *Parameter) error
}

// 参数.
type Parameter struct {
	Token string                `form:"token" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
	AppID string                `form:"app_id" binding:"required"`

	Universal
}

type Universal struct {
	// XXX: xxxxxxxx
	UploadFilePath string
	CssFilePath    string

	StorePath string
}
