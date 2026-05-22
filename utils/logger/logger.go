package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"text/template"
	"time"
)

type Environment string

const (
	EnvironmentProduction  Environment = "prod"
	EnvironmentDevelopment Environment = "dev"
)

type Level uint32

const (
	LevelDefault Level = iota
	LevelFatal
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
	LevelTrace

	LevelCritical = LevelFatal

	NullStr          = "null"
	DefaultCallDepth = 2
)

func (l Level) String() string {
	switch l {
	case LevelDefault:
		return "Default"

	case LevelFatal:
		return "CRITICAL"

	case LevelError:
		return "ERROR"

	case LevelWarning:
		return "WARNING"

	case LevelInfo:
		return "INFO"

	case LevelDebug:
		return "DEBUG"

	case LevelTrace:
		return "TRACE"
	}

	return "Unknown"
}

func StringLevel(s string) Level {
	switch strings.ToLower(s) {
	case "default":
		return LevelDefault

	case "critical":
		return LevelCritical

	case "fatal":
		return LevelFatal

	case "error":
		return LevelError

	case "warning":
		return LevelWarning

	case "info":
		return LevelInfo

	case "debug":
		return LevelDebug

	case "trace":
		return LevelTrace
	}

	return LevelDefault
}

type Context struct {
	wtr  io.Writer
	tmpl *template.Template
	lvl  Level
}

var (
	DefaultFormat    = "{{timeStamp .Now}}[{{.Level}}][NULL, {{.Function}}({{.Filename}}:{{.LineNo}})][{{.Trigger}}]: {{.Message}} "
	DefaultContext   = NewContext(os.Stderr, DefaultFormat, LevelWarning)
	DefaultV1Format  = `{"timestamp":"{{iso8601TimeStamp .Now}}", "triggerLabel":{{.TriggerLabel}}, "logLevel":"{{.Level}}", "logFacility":"{{.LogFacility}}", "threadId":{{.ThreadId}}, "function":"{{.Function}}", "file":"{{.Filename}}", "lineNo":{{.LineNo}}, "message":"{{.Message}}", "extra": {{.Extra}}}`
	DefaultV1Context = NewContext(os.Stderr, DefaultV1Format, LevelWarning)

	// Keep track of log file for proper closure
	logFile  *os.File
	logMutex sync.Mutex
)

type XFields map[string]interface{}

type Logger struct {
	*Context
	subfacility string
	lvl         Level
	buf         []byte
	mux         sync.Mutex
	calldepth   int
}

func timeStamp(t time.Time) string {
	return t.String()
}

// Returns the timestamp of iso8601 format, ex : 2023-04-26T08:45:05.232Z
func iso8601TimeStamp(t time.Time) string {
	// Timestamp of format : 2023-04-26T08:45:05.232Z
	return t.Format("2006-01-02T15:04:05.000Z")
}

var (
	packageInitializers []func()
	initMutex           sync.Mutex
	loggerInitialized   bool
)

// RegisterPackageLogger registers a function to initialize package loggers
// This should be called in each package's init() function
func RegisterPackageLogger(initFunc func()) {
	initMutex.Lock()
	defer initMutex.Unlock()

	if loggerInitialized {
		// If logger is already initialized, run immediately
		initFunc()
	} else {
		// Otherwise, queue it for later
		packageInitializers = append(packageInitializers, initFunc)
	}
}

// InitializeAllLoggers sets up logging and initializes all registered package loggers
func InitializeAllLoggers() error {
	initMutex.Lock()
	defer initMutex.Unlock()

	if loggerInitialized {
		return nil // Already initialized
	}

	// Set up file logging
	if err := InitLogger(); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	// Set global log level
	DefaultV1Context.SetLevel(LevelTrace)

	// Initialize all registered package loggers
	for _, initFunc := range packageInitializers {
		initFunc()
	}

	loggerInitialized = true

	// Clear the initializers slice to free memory
	packageInitializers = nil

	return nil
}

// InitLogger sets up logging to both console and logs/app.log
func InitLogger() error {
	logMutex.Lock()
	defer logMutex.Unlock()

	// Close existing log file if open
	if logFile != nil {
		logFile.Close()
	}

	// Ensure logs directory exists
	if err := os.MkdirAll("logs", 0755); err != nil {
		return fmt.Errorf("failed to create logs directory: %w", err)
	}

	// Open log file (append mode)
	var err error
	logFile, err = os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	// Create MultiWriter to write to both console and file
	multiWriter := io.MultiWriter(os.Stderr, logFile)

	// Redirect contexts to write to both console and file
	DefaultContext = NewContext(multiWriter, DefaultFormat, LevelWarning)
	DefaultV1Context = NewContext(multiWriter, DefaultV1Format, LevelWarning)

	return nil
}

// CloseLogger properly closes the log file
func CloseLogger() error {
	logMutex.Lock()
	defer logMutex.Unlock()

	if logFile != nil {
		err := logFile.Close()
		logFile = nil
		return err
	}
	return nil
}

func NewContext(w io.Writer, fmt string, lvl Level) *Context {
	tmpl := template.New("LogPrefix")
	tmpl = tmpl.Funcs(map[string]interface{}{"timeStamp": timeStamp, "iso8601TimeStamp": iso8601TimeStamp})
	tmpl = template.Must(tmpl.Parse(fmt))
	return &Context{wtr: w,
		tmpl: tmpl,
		lvl:  lvl}
}

func (ctx *Context) SetLevel(lvl Level) {
	atomic.StoreUint32((*uint32)(&ctx.lvl), uint32(lvl))
}

func (ctx *Context) Level() Level {
	return Level(atomic.LoadUint32((*uint32)(&ctx.lvl)))
}

// Get the logger context
func (ctx *Context) GetLogger(sub string, lvl Level) *Logger {
	return &Logger{Context: ctx,
		subfacility: sub,
		lvl:         lvl,
		mux:         sync.Mutex{},
		calldepth:   DefaultCallDepth,
	}
}

func (l *Logger) SetLevel(lvl Level) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.lvl = lvl
}

func (l *Logger) Level() Level {
	l.mux.Lock()
	defer l.mux.Unlock()
	return l.lvl
}

func (l *Logger) LogFacility() string {
	l.mux.Lock()
	defer l.mux.Unlock()

	return l.subfacility
}

// CallDepth is used for logging calling function's name, line number and file name.
// Default CallDepth is 2, if using a wrapper set this value to 3
func (l *Logger) SetCallDepth(callDepth int) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.calldepth = callDepth
}

func (l *Logger) Output(calldepth int, lvl Level, triggerLabel string, m string) error {
	now := time.Now()
	pc, filename, lineno, _ := runtime.Caller(calldepth)
	f := runtime.FuncForPC(pc)

	if triggerLabel == "" {
		triggerLabel = NullStr
	}

	//Variable replacement
	vars := map[string]interface{}{"Now": now,
		"Level":        lvl,
		"Function":     f.Name(),
		"Filename":     filepath.Base(filename),
		"LineNo":       lineno,
		"TriggerLabel": triggerLabel,
		"ThreadId":     os.Getpid(),
		"Message":      m}

	//Execute template
	l.mux.Lock()
	defer l.mux.Unlock()

	var tbuf bytes.Buffer

	l.buf = l.buf[:0]

	if err := l.tmpl.ExecuteTemplate(&tbuf, "LogPrefix", vars); err != nil {
		return err
	}

	l.buf = append(l.buf, tbuf.Bytes()...)
	if len(l.buf) > 0 && l.buf[len(l.buf)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}

	_, err := l.wtr.Write(l.buf)

	// Force flush to ensure data is written to disk
	if syncer, ok := l.wtr.(interface{ Sync() error }); ok {
		syncer.Sync()
	}

	return err
}

/*
OutputX is similar to Output with additional support for extra logging fields.
The `XFields` will be json encoded for logging. Use {{.Extra}} in formatting template without any surrounding quotes.
Example:

	logger.OutputX(2,"INFO","",&logging.XFields{"key1":"StrValue","key2":IntValue},"This is an example")
*/
func (l *Logger) OutputX(calldepth int, lvl Level, triggerLabel string, xFields *XFields, m string) error {
	now := time.Now()
	pc, filename, lineno, _ := runtime.Caller(calldepth)
	f := runtime.FuncForPC(pc)

	if triggerLabel == "" {
		triggerLabel = NullStr
	}

	//Variable replacement
	vars := map[string]interface{}{"Now": now,
		"Level":        lvl,
		"Function":     f.Name(),
		"Filename":     filepath.Base(filename),
		"LineNo":       lineno,
		"LogFacility":  l.subfacility,
		"TriggerLabel": triggerLabel,
		"ThreadId":     os.Getpid(),
		"Message":      m}

	// assigning null if no xFields or any marshal error
	if xFields != nil && len(*xFields) > 0 {
		extraByte, marshalErr := json.Marshal(xFields)
		if marshalErr != nil {
			vars["Extra"] = NullStr
		} else {
			vars["Extra"] = string(extraByte)
		}
	} else {
		vars["Extra"] = NullStr
	}

	//Execute template
	l.mux.Lock()
	defer l.mux.Unlock()
	var tbuf bytes.Buffer
	l.buf = l.buf[:0]
	if err := l.tmpl.ExecuteTemplate(&tbuf, "LogPrefix", vars); err != nil {
		return err
	}

	l.buf = append(l.buf, tbuf.Bytes()...)
	if len(l.buf) > 0 && l.buf[len(l.buf)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}

	_, err := l.wtr.Write(l.buf)

	// Force flush to ensure data is written to disk
	if syncer, ok := l.wtr.(interface{ Sync() error }); ok {
		syncer.Sync()
	}

	return err
}

func (l *Logger) CheckLevel(lvl Level) bool {
	llvl := l.lvl
	if llvl == LevelDefault {
		llvl = l.Context.Level()
	}

	return llvl >= lvl
}

func (l *Logger) Log(lvl Level, v ...interface{}) {
	if lvl == LevelFatal || l.CheckLevel(lvl) {
		l.Output(l.calldepth, lvl, "", fmt.Sprint(v...))
		if lvl == LevelFatal {
			panic(fmt.Sprint(v...))
		}
	}
}

func (l *Logger) Logf(lvl Level, str string, v ...interface{}) {
	if lvl == LevelFatal || l.CheckLevel(lvl) {
		l.Output(l.calldepth, lvl, "", fmt.Sprintf(str, v...))
		if lvl == LevelFatal {
			panic(fmt.Sprintf(str, v...))
		}
	}
}

func (l *Logger) Logxf(lvl Level, xFields *XFields, str string, v ...interface{}) {
	if lvl == LevelFatal || l.CheckLevel(lvl) {
		l.OutputX(l.calldepth, lvl, "", xFields, fmt.Sprintf(str, v...))

		if lvl == LevelFatal {
			panic(str)
		}
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.CheckLevel(LevelError) {
		l.Output(l.calldepth, LevelError, "", fmt.Sprint(v...))
	}
}

func (l *Logger) Errorf(str string, v ...interface{}) {
	if l.CheckLevel(LevelError) {
		l.Output(l.calldepth, LevelError, "", fmt.Sprintf(str, v...))
	}
}

// logger.ErrorX(&log.XFields{"StrField": "str", "IntField": 2, "floatField": 3.2}, "Example log")
func (l *Logger) Errorxf(xFields *XFields, str string, v ...interface{}) {
	if l.CheckLevel(LevelError) {
		l.OutputX(l.calldepth, LevelError, "", xFields, fmt.Sprintf(str, v...))
	}
}

func (l *Logger) Warning(v ...interface{}) {
	if l.CheckLevel(LevelWarning) {
		l.Output(l.calldepth, LevelWarning, "", fmt.Sprint(v...))
	}
}

func (l *Logger) Warningf(str string, v ...interface{}) {
	if l.CheckLevel(LevelWarning) {
		l.Output(l.calldepth, LevelWarning, "", fmt.Sprintf(str, v...))
	}
}

// Example:
//
//	logger.Warningxf(&log.XFields{"StrField": "str", "IntField": 2, "floatField": 3.2}, "Example log")
func (l *Logger) Warningxf(xFields *XFields, str string, v ...interface{}) {
	if l.CheckLevel(LevelWarning) {
		l.OutputX(l.calldepth, LevelWarning, "", xFields, fmt.Sprintf(str, v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.CheckLevel(LevelInfo) {
		l.Output(l.calldepth, LevelInfo, "", fmt.Sprint(v...))
	}
}

func (l *Logger) Infof(str string, v ...interface{}) {
	if l.CheckLevel(LevelInfo) {
		l.Output(l.calldepth, LevelInfo, "", fmt.Sprintf(str, v...))
	}
}

// Example:
//
//	logger.Infoxf(&log.XFields{"StrField": "str", "IntField": 2, "floatField": 3.2}, "Example log")
func (l *Logger) Infoxf(xFields *XFields, str string, v ...interface{}) {
	if l.CheckLevel(LevelInfo) {
		l.OutputX(l.calldepth, LevelInfo, "", xFields, fmt.Sprintf(str, v...))
	}
}

func (l *Logger) Debug(v ...interface{}) {
	if l.CheckLevel(LevelDebug) {
		l.Output(l.calldepth, LevelDebug, "", fmt.Sprint(v...))
	}
}

func (l *Logger) Debugf(str string, v ...interface{}) {
	if l.CheckLevel(LevelDebug) {
		l.Output(l.calldepth, LevelDebug, "", fmt.Sprintf(str, v...))
	}
}

// Example:
//
//	logger.Debugxf(&log.XFields{"StrField": "str", "IntField": 2, "floatField": 3.2}, "Example log")
func (l *Logger) Debugxf(xFields *XFields, str string, v ...interface{}) {
	if l.CheckLevel(LevelDebug) {
		l.OutputX(l.calldepth, LevelDebug, "", xFields, fmt.Sprintf(str, v...))
	}
}

func (l *Logger) Trace(v ...interface{}) {
	if l.CheckLevel(LevelTrace) {
		l.Output(l.calldepth, LevelTrace, "", fmt.Sprint(v...))
	}
}

func (l *Logger) Tracef(str string, v ...interface{}) {
	if l.CheckLevel(LevelTrace) {
		l.Output(l.calldepth, LevelTrace, "", fmt.Sprintf(str, v...))
	}
}

func (l *Logger) Tracexf(xFields *XFields, str string, v ...interface{}) {
	if l.CheckLevel(LevelTrace) {
		l.OutputX(l.calldepth, LevelTrace, "", xFields, fmt.Sprintf(str, v...))
	}
}

func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.Output(l.calldepth, LevelFatal, "", s)
	panic(s)
}

func (l *Logger) Panicf(str string, v ...interface{}) {
	s := fmt.Sprintf(str, v...)
	l.Output(l.calldepth, LevelFatal, "", s)
	panic(s)
}

func (l *Logger) Panicxf(xFields *XFields, str string, v ...interface{}) {
	s := fmt.Sprintf(str, v...)
	l.OutputX(l.calldepth, LevelTrace, "", xFields, s)
	panic(s)
}
