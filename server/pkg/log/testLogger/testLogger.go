package testLogger

import (
	"fmt"
	"net/http"
	"os"

	"github.com/darwinfroese/hacksite/server/pkg/log"
)

type testLogger struct {
}

// New creates a new test logger
func New() log.Logger {
	return &testLogger{}
}

func (l *testLogger) Info(message string) {
	fmt.Fprintf(os.Stdout, "%s\n", message)
}

func (l *testLogger) Error(message string) {
	fmt.Fprintf(os.Stderr, "%s\n", message)
}

func (l *testLogger) InfoWithRequest(r *http.Request, requestID, message string) {
	fmt.Fprintf(os.Stderr, "%s\n", message)
}

func (l *testLogger) ErrorWithRequest(r *http.Request, requestID, message string) {
	fmt.Fprintf(os.Stderr, "%s\n", message)
}
