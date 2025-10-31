package logger

import (
	"github.com/szwtdl/chinatax-system/internal/core/types"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func GetLogger(cfg types.LogOptions) *zap.SugaredLogger {
	if sugarLogger != nil {
		return sugarLogger
	}
	var level zapcore.Level
	mode := strings.ToLower(cfg.Mode)
	if mode == "production" {
		level = zapcore.ErrorLevel
	} else {
		level = zapcore.DebugLevel
	}
	logFile := cfg.FilePath
	if logFile == "" {
		logFile = filepath.Join("logs", "app-"+mode+".log")
	} else {
		dir := filepath.Dir(logFile)
		base := filepath.Base(logFile)
		ext := filepath.Ext(base)
		name := strings.TrimSuffix(base, ext)
		base = name + "-" + mode + ext // 在文件名前加环境标识
		logFile = filepath.Join(dir, base)
	}
	encoder := getEncoder()
	writerSyncer := getLogWriter(logFile, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	fileCore := zapcore.NewCore(encoder, writerSyncer, level)
	consoleOutput := zapcore.Lock(os.Stdout)
	consoleCore := zapcore.NewCore(encoder, consoleOutput, level)
	core := zapcore.NewTee(fileCore, consoleCore)
	logger = zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
	return sugarLogger
}

func getEncoder() zapcore.Encoder {
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(FileName string, MaxSize, MaxBackups, MaxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   FileName,
		MaxSize:    MaxSize, // MB
		MaxBackups: MaxBackups,
		MaxAge:     MaxAge, // days
		Compress:   true,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
