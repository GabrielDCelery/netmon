package netstat

import (
	"context"
	"runtime"
	"testing"
)

func TestSSRunner_Run(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("ss command is only available on Linux")
	}

	runner := NewSSRunner()
	output, err := runner.Run(context.Background())
	if err != nil {
		t.Fatalf("SSRunner.Run() error: %v", err)
	}

	if len(output) == 0 {
		t.Error("SSRunner.Run() returned empty output")
	}
}
