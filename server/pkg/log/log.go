package log

import "net/http"

// Logger is a simple logging interface for outputting
// messages to log files
type Logger interface {
	Info(string)
	Error(string)

	ErrorWithRequest(request *http.Request, requestID string, message string)
	InfoWithRequest(request *http.Request, requestID string, message string)
}
