package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/quanxiang-cloud/fileserver/pkg/guide"
)

var (
	filePath  = flag.String("filePath", "", "file path")
	storePath = flag.String("storePath", "", "store path")
)

func main() {
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		panic(err)
	}

	fi, err := file.Stat()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		panic(err)
	}

	g, err := guide.NewGuide()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		panic(err)
	}

	os.Stdout.WriteString("uploading...")

	os.Stdout.WriteString(fmt.Sprintf("size: %d", fi.Size()))
	os.Stdout.WriteString(fmt.Sprintf("size: %+v", file))

	err = g.FutileUploadFile(context.Background(), *storePath, file, fi.Size(), guide.Readable)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		panic(err)
	}

	os.Stdout.WriteString("done")
}
