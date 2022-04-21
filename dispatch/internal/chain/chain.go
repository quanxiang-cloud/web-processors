package chain

import (
	"context"
	"fmt"
	"io"
	"os/exec"
)

// Command is the interface that wraps the basic 'Do' method.
type Command interface {
	Do(ctx context.Context, args *Parameter) error
}

// Parameter parameter.
type Parameter struct {
	AppID   string `json:"appID"`
	Element string `json:"element" binding:"required"`

	Universal
}

// Universal is the universal parameter.
type Universal struct {
	UploadFilePath string
	CssFilePath    string
	CssFileHash    string
	StorePath      string
}

const (
	uploadFileName = "input_file.txt"
)

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
	InputFile       = "INPUT_FILE"
	PersonaHostName = "PERSONA_HOSTNAME"
)

func genEnv(key, value string) string {
	return fmt.Sprintf("%s=%s", key, value)
}

func execute(cmd *exec.Cmd) (string, error) {
	// get stdout and stderr readCloser
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	defer stdout.Close()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}
	defer stderr.Close()

	// start command.
	if err := cmd.Start(); err != nil {
		return "", err
	}

	stdoutBuf, err := io.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	stderrBuf, err := io.ReadAll(stderr)
	if err != nil {
		return "", err
	}

	// wait for the command to complete
	// nolint: errcheck
	cmd.Wait()

	if !cmd.ProcessState.Success() {
		return "", fmt.Errorf("%s", stderrBuf)
	}

	return string(stdoutBuf), nil
}
