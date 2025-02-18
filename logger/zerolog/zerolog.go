package zerolog

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"osh.com/rps/registrar/config"
)

type Logger struct {
	logger  *zerolog.Logger
	mu      *sync.Mutex
	rootDir string
}

func GetNewLogger(loggerConf *config.LoggerConfig, institute string) Logger {

	return Logger{
		logger:  loadLogger(loggerConf, institute),
		mu:      &sync.Mutex{},
		rootDir: loggerConf.Logger.RootDir,
	}
}

func loadLogger(conf *config.LoggerConfig, institute string) *zerolog.Logger {

	var logLevel zerolog.Level

	file := ConfigLumberjack(conf, institute)

	switch conf.Logger.LogLevel {
	case "disable":
		logLevel = zerolog.Disabled
	case "trace":
		logLevel = zerolog.TraceLevel
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	case "warn":
		logLevel = zerolog.WarnLevel
	case "fatal":
		logLevel = zerolog.FatalLevel
	case "panic":
		logLevel = zerolog.PanicLevel
	default:
		logLevel = zerolog.GlobalLevel()
	}

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	multiLogger := zerolog.MultiLevelWriter(consoleWriter, file)
	new := zerolog.New(multiLogger).
		Level(logLevel).
		With().
		Timestamp().
		Logger()

	return &new
}

func ConfigLumberjack(conf *config.LoggerConfig, institute string) *lumberjack.Logger {

	var fullPath, LogDir string

	logFileName := conf.Logger.LogFile
	logFilePath := conf.Logger.LogPath
	maxSize := conf.Logger.MaxSize
	maxBackups := conf.Logger.MaxBackups
	maxAge := conf.Logger.MaxAge
	compress := conf.Logger.Compress

	if institute == "" {
		LogDir = conf.Logger.DefaultLogDir
	} else {
		LogDir = institute
	}

	if logFilePath != "" || LogDir != "" || logFileName != "" {
		fullPath = fmt.Sprintf("%s/%s/%s", logFilePath, LogDir, logFileName)
	} else {
		fullPath = "log/pos_service.log"
	}

	return &lumberjack.Logger{
		Filename:   fullPath,
		MaxSize:    maxSize,    // megabytes
		MaxBackups: maxBackups, // number of backups
		MaxAge:     maxAge,     // days
		Compress:   compress,   // compress backups
	}

}

func fileAndLine(skip int, rootDir string) (file string, line int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown path", 0
	}
	if strings.Contains(file, "/"+rootDir+"/") {
		file = file[strings.Index(file, "/"+rootDir+"/"):]
	}
	return
}

func (l Logger) Info(messages ...string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	message := strings.Join(messages, " ")

	logger := l.logger

	file, line := fileAndLine(2, l.rootDir)
	logger.Info().Msgf("%s:%d "+message, file, line)
}

func (l Logger) Warn(messages ...string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	message := strings.Join(messages, " ")

	logger := l.logger

	file, line := fileAndLine(2, l.rootDir)
	logger.Warn().Msgf("%s:%d "+message, file, line)

}

func (l Logger) Error(messages ...string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	message := strings.Join(messages, " ")

	logger := l.logger

	file, line := fileAndLine(2, l.rootDir)
	logger.Error().Msgf("%s:%d "+message, file, line)
}

func (l Logger) Fatal(messages ...string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	message := strings.Join(messages, " ")

	logger := l.logger

	file, line := fileAndLine(2, l.rootDir)
	logger.Fatal().Msgf("%s:%d "+message, file, line)
}

func (l Logger) Debug(messages ...string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	message := strings.Join(messages, " ")

	logger := l.logger

	file, line := fileAndLine(2, l.rootDir)
	logger.Debug().Msgf("%s:%d "+message, file, line)
}

func (l Logger) Trace(messages ...string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	message := strings.Join(messages, " ")

	logger := l.logger

	file, line := fileAndLine(2, l.rootDir)
	logger.Trace().Msgf("%s:%d "+message, file, line)
}

func (l Logger) Panic(messages ...string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	message := strings.Join(messages, " ")

	logger := l.logger

	file, line := fileAndLine(2, l.rootDir)
	logger.Panic().Msgf("%s:%d "+message, file, line)
}

func (l Logger) PanicRecovery(messages ...string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	message := strings.Join(messages, " ")

	logger := l.logger

	file, line := fileAndLine(3, l.rootDir)
	logger.Error().Msgf("%s:%d "+message, file, line)
	logger.Info().Msg("Recovered from panic.")
}

// func (l Logger) RequestLogger(messages map[string]string) {
// 	l.mu.Lock()
// 	defer l.mu.Unlock()

// 	logger := l.logger

// 	logCxt := logger.With()

// 	for key, val := range messages {
// 		if val != "" {
// 			logCxt = logCxt.Str(key, val)
// 		}
// 	}

// 	logger = logCxt.Logger()
// 	logger = updatedLogger
// 	logger.Info().Msg("")

// 	// if s, ok := messages["Status"]; ok && s == "Success" {
// 	// 	logger.Info().Msg("")
// 	// } else {
// 	// 	logger.Error().Msg("")
// 	// }

// }
