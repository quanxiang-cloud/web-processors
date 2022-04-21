package code

import error2 "github.com/quanxiang-cloud/cabin/error"

func init() {
	error2.CodeTable = CodeTable
}

// error code.
const (
	ErrSaveFile = 99001500001
	ErrExecute  = 99002400001
)

// CodeTable is the code table.
var CodeTable = map[int64]string{
	ErrSaveFile: "save file error",
	ErrExecute:  "execute task error",
}
