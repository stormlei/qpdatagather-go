package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var sugarLogger *zap.SugaredLogger

func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}
func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}

func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}
func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}

func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}
func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}

func InitLogger(conf *viper.Viper) {
	writeSyncer := getLogWriter()
	encoder := getEncoder()

	logLevel := zapcore.ErrorLevel
	logLevelConf := conf.GetString("log.level")
	switch logLevelConf {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	default:
		logLevel = zapcore.ErrorLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, logLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "logs/logs.log",
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func Sync() error {
	return sugarLogger.Sync()
}
