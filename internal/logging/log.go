package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/max-gryshin/stskp_sport_api/internal/file"
	"github.com/max-gryshin/stskp_sport_api/internal/setting"
)

type Level int

var (
	f *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Setup initialize the log instance
func Setup(config *setting.App) {
	var err error
	filePath := filepath.Join(config.RuntimeRootPath, config.LogSavePath)
	fileName := strings.Join([]string{config.LogSaveName, config.LogFileExt}, ".")
	f, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(f, DefaultPrefix, log.LstdFlags)
}

// Close closes output file
func Close() error {
	if f != nil {
		return f.Close()
	}

	return nil
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, fileName, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(fileName), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
