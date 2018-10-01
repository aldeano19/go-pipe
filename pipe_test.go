package pipe

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

func TestCatAndGrep(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("skipping test, OS is not linux. ")
	}

	filePath := "/home/erieljr1/go/src/github.com/aldeano19/go-pipe/pipe_test.go"

	testGrepLine := "Test Grep Line"

	var b bytes.Buffer
	Command(&b,
		exec.Command("cat", filePath),
		exec.Command("grep", testGrepLine),
	)

	actual := b.String()

	if !strings.Contains(actual, testGrepLine) {
		t.Errorf("expected: %s, got: %s", testGrepLine, actual)
	}
}
