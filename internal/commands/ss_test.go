package commands

import (
	"context"
	"runtime"
	"testing"
)

func TestSSCommand_Run(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("ss command is only available on Linux")
	}

	runner := NewSSCommand()
	output, err := runner.Run(context.Background())
	if err != nil {
		t.Fatalf("SSCommand.Run() error: %v", err)
	}

	if len(output) == 0 {
		t.Error("SSCommand.Run() returned empty output")
	}
}
