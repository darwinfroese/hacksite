package logrus

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/darwinfroese/hacksite/server/pkg/log"
	"github.com/sirupsen/logrus"
)

var fileLock = &sync.Mutex{}

type logrusLogger struct {
	logger      *logrus.Logger
	logLocation string
}

// New returns a new logrus logger
func New(location string) log.Logger {
	return &logrusLogger{
		logger:      logrus.New(),
		logLocation: location,
	}
}

func (l *logrusLogger) ErrorWithRequest(r *http.Request, id, message string) {
	fileLock.Lock()

	f := getFile(l.logLocation)
	l.logger.Out = f
	defer func() {
		f.Close()
		fileLock.Unlock()
	}()

	name, file, line := getCallingFunction()

	l.logger.WithFields(logrus.Fields{
		"request-id":   id,
		"request-url":  r.URL.Path,
		"user-agent":   r.UserAgent(),
		"user-address": r.RemoteAddr,
		"calling-func": name,
		"calling-file": file,
		"calling-line": line,
	}).Error(message)
}

func (l *logrusLogger) InfoWithRequest(r *http.Request, id, message string) {
	fileLock.Lock()

	f := getFile(l.logLocation)
	l.logger.Out = f
	defer func() {
		f.Close()
		fileLock.Unlock()
	}()

	name, file, line := getCallingFunction()

	l.logger.WithFields(logrus.Fields{
		"request-id":   id,
		"request-url":  r.URL.Path,
		"user-agent":   r.UserAgent(),
		"user-address": r.RemoteAddr,
		"calling-func": name,
		"calling-file": file,
		"calling-line": line,
	}).Info(message)
}

func (l *logrusLogger) Info(message string) {
	fileLock.Lock()

	f := getFile(l.logLocation)
	l.logger.Out = f
	defer func() {
		f.Close()
		fileLock.Unlock()
	}()

	name, file, line := getCallingFunction()

	l.logger.WithFields(logrus.Fields{
		"calling-func": name,
		"calling-file": file,
		"calling-line": line,
	}).Info(message)
}

func (l *logrusLogger) Error(message string) {
	fileLock.Lock()

	f := getFile(l.logLocation)
	l.logger.Out = f
	defer func() {
		f.Close()
		fileLock.Unlock()
	}()

	name, file, line := getCallingFunction()

	l.logger.WithFields(logrus.Fields{
		"calling-func": name,
		"calling-file": file,
		"calling-line": line,
	}).Error(message)
}

func getFile(location string) *os.File {
	d := time.Now().Format("02012006")
	fn := fmt.Sprintf("%s/%s.log", location, d)

	file, err := os.OpenFile(fn, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return os.Stderr
	}

	return file
}

func getCallingFunction() (string, string, int) {
	// we get the callers as uintptrs - but we just need 1
	fpcs := make([]uintptr, 1)

	// skip 3 levels to get to the caller of whoever called Caller()
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return "unknown", "unknown", 0
	}

	// get the info of the actual function that's in the pointer
	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return "unknown", "unknown", 0
	}

	// return its name
	file, line := fun.FileLine(fpcs[0] - 1)
	return fun.Name(), file, line
}
