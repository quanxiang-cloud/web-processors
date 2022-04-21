package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/quanxiang-cloud/cabin/logger"
	router "github.com/quanxiang-cloud/web-processors/dispatch/api"
	"github.com/quanxiang-cloud/web-processors/dispatch/pkg/config"
)

var (
	port            string
	uploadDir       string
	commandDir      string
	storePathPrefix string

	personaEndpoint  string
	personaKeySuffix string
	personaVersion   string

	level           int
	development     bool
	initial         int
	thereafter      int
	outputPath      Arr
	errorOutputPath Arr
)

func main() {
	loggerFlag()
	logger.Logger = logger.New(&logger.Config{
		Level:       level,
		Development: development,
		Sampling: logger.Sampling{
			Initial:    initial,
			Thereafter: thereafter,
		},
		OutputPath:      outputPath,
		ErrorOutputPath: errorOutputPath,
	})
	conf := &config.Config{
		Port:            port,
		UploadDir:       uploadDir,
		CommandDir:      commandDir,
		StorePathPrefix: storePathPrefix,

		PersonaEndpoint:  personaEndpoint,
		PersonaKeySuffix: personaKeySuffix,
		PersonaVersion:   personaVersion,
	}
	router, err := router.NewRouter(conf)
	if err != nil {
		panic(err)
	}

	go router.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

type Arr []string

func (a *Arr) String() string {
	return fmt.Sprint(*a)
}

func (a *Arr) Set(s string) error {
	*a = append(*a, strings.Split(s, ",")...)
	return nil
}

func loggerFlag() {
	flag.StringVar(&port, "port", ":80", "port")
	flag.StringVar(&uploadDir, "upload-dir", ".", "upload dir")
	// NOTE: command save directory is consistent with the save directory in dockerfile
	flag.StringVar(&commandDir, "command-dir", "scripts", "command dir")
	flag.StringVar(&storePathPrefix, "store-path-prefix", "web-processors/config/css", "store path prefix")
	flag.StringVar(&personaEndpoint, "persona-endpoint", "http://persona", "persona endpoint")
	flag.StringVar(&personaKeySuffix, "persona-key-suffix", "style_guide_css:draft", "persona key suffix")
	flag.StringVar(&personaVersion, "persona-version", "1.0.0", "persona version")

	flag.IntVar(&level, "level", -1, "log level")
	flag.BoolVar(&development, "development", false, "log development")
	flag.IntVar(&initial, "initial", 100, "log initial")
	flag.IntVar(&thereafter, "thereafter", 100, "log thereafter")
	flag.Var(&outputPath, "outputPath", "log outputPath")
	flag.Var(&errorOutputPath, "errorOutputPath", "log errorOutputPath")

	flag.Parse()

	if len(outputPath) == 0 {
		outputPath.Set("stderr")
	}

	if len(errorOutputPath) == 0 {
		errorOutputPath.Set("stderr")
	}
}
