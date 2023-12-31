package logging

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"greatestworks/aop/colors"
	"greatestworks/aop/logtype"
	"greatestworks/aop/protos"
)

type Options struct {
	App        string // Service Weaver application (e.g., "todo")
	Deployment string // Service Weaver deployment (e.g., "36105c89-85b1...")
	Component  string // Service Weaver component (e.g., "Todo")
	Weavelet   string // Service Weaver weavelet id (e.g., "36105c89-85b1...")

	Attrs []string
}

// makeEntry returns an entry that is fully populated with the provided level,
// msg, and attributes. The file and line information is taken from the call
// stack after skipping frameskip entries up. For example, a frameskip of 0
// will use the file and line of the call to makeEntry, and a frameskip of 1
// will use the file and line of the caller's parent.
func makeEntry(level, msg string, attrs []any, frameskip int, opts Options) *protos.LogEntry {
	// TODO(sanjay): Is it necessary to copy opts.Attrs even if no new attrs
	// are being added?
	entry := protos.LogEntry{
		App:        opts.App,
		Version:    opts.Deployment,
		Component:  opts.Component,
		Node:       opts.Weavelet,
		TimeMicros: time.Now().UnixMicro(),
		Level:      level,
		File:       "",
		Line:       -1,
		Msg:        msg,
		Attrs:      logtype.AppendAttrs(opts.Attrs, attrs),
	}

	// We add one to frameskip because we also skip makeEntry's frame.
	_, file, line, ok := runtime.Caller(frameskip + 1)
	if ok {
		entry.File = file
		entry.Line = int32(line)
	}

	return &entry
}

// FuncLogger is a logger that calls a supplied function on every log entry.
type FuncLogger struct {
	Opts  Options                      // configures the log entries
	Write func(entry *protos.LogEntry) // called on every log entry
}

var _ logtype.Logger = FuncLogger{}

// Debug implements the [weaver.Logger] interface.
func (l FuncLogger) Debug(msg string, attrs ...any) {
	l.Write(makeEntry("debug", msg, attrs, 1, l.Opts))
}

// Info implements the [weaver.Logger] interface.
func (l FuncLogger) Info(msg string, attrs ...any) {
	l.Write(makeEntry("info", msg, attrs, 1, l.Opts))
}

// Error implements the [weaver.Logger] interface.
func (l FuncLogger) Error(msg string, err error, attrs ...any) {
	e := makeEntry("error", msg, attrs, 1, l.Opts)
	if err != nil {
		e.Attrs = append(e.Attrs, "err", err.Error())
	}
	l.Write(e)
}

// StderrLogger returns a logger that pretty prints log entries to stderr.
func StderrLogger(opts Options) FuncLogger {
	pp := NewPrettyPrinter(colors.Enabled())
	writeText := func(entry *protos.LogEntry) {
		fmt.Fprintln(os.Stderr, pp.Format(entry))
	}
	return FuncLogger{opts, writeText}
}

// TestLogger pretty prints log entries using t.Log.
type TestLogger struct {
	FuncLogger                // underlying weaver.Logger
	t          testing.TB     // logs until t finishes
	pp         *PrettyPrinter // pretty prints log entries
	mu         sync.Mutex     // guards finished
	finished   bool           // has t finished?
}

// NewTestLogger returns a new TestLogger.
func NewTestLogger(t testing.TB) *TestLogger {
	logger := &TestLogger{t: t, pp: NewPrettyPrinter(colors.Enabled())}
	t.Cleanup(logger.Silence)
	logger.FuncLogger = FuncLogger{
		Opts:  Options{Component: "TestLogger", Weavelet: uuid.New().String()},
		Write: logger.Log,
	}
	return logger
}

// Log logs the provided log entry using t.Log.
func (t *TestLogger) Log(entry *protos.LogEntry) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.finished {
		// If the test is finished, Log may panic if called, so don't log
		// anything.
		return
	}
	entry.TimeMicros = time.Now().UnixMicro()
	t.t.Log(t.pp.Format(entry))
}

// Silence prevents any future log entries from being logged.
func (t *TestLogger) Silence() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.finished = true
}
