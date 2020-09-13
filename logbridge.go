// Package logbridge implements an io.Writer proxy for log.Logger
// suitable for use as an output for another log.Logger.
package logbridge

import (
	"io"
	"log"
	"strings"
)

var _ io.Writer = LogBridge{nil}

type LogBridge struct {
	*log.Logger
}

// Wrap returns a new log.Logger that writes to the given logger
// through the LogBridge. The flags are usually unset (zero).
func Wrap(l *log.Logger, prefix string, flags int) *log.Logger {
	return log.New(LogBridge{l}, prefix, flags)
}

func (lb LogBridge) Write(p []byte) (int, error) {
	var b strings.Builder
	b.Write(p)

	// We’ll be called from log.Logger
	// Print, Fatal and Panic functions.
	//
	// The call stack goes as follows:
	//
	//   1. log.Logger                 Print(message)
	//   2. log.Logger                 Output(2, message)
	//   3. logbridge.LogBridge        Write(message)
	//   4. logbridge.LogBridge.Logger Output(4, message)
	//
	err := lb.Output(4, b.String())

	return len(p), err
}
