package main

import (
	"context"
	"flag"
	"os"

	"github.com/quanxiang-cloud/fileserver/pkg/guide"
)

var (
	filePath   = flag.String("filePath", "", "file path")
	uploadPath = flag.String("uploadPath", "", "upload path")
)

func main() {
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.Sync()

		return
	}

	fi, err := file.Stat()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.Sync()

		return
	}

	g, err := guide.NewGuide()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.Sync()

		return
	}

	err = g.FutileUploadFile(context.Background(), *uploadPath, file, fi.Size(), guide.Readable)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.Sync()

		return
	}
}
