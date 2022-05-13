package main

import (
	"context"
	"flag"
	"log"
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
		panic(err)
	}

	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}

	g, err := guide.NewGuide()
	if err != nil {
		panic(err)
	}

	log.Println("uploading...")

	err = g.FutileUploadFile(context.Background(), *storePath, file, fi.Size(), guide.Readable)
	if err != nil {
		panic(err)
	}

	log.Println("Done...")
}
