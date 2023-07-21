package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// var logger *zap.SugaredLogger
// var logFile string = "./test.log"
// var logErrFile string = "./test.err.log"

// 使用json格式打印日志
// 日志输出到标准输出和日志文件中
// 错误日志单独输出一份到标准错误和错误日志文件中
// 以上日志文件均根据文件大小轮转
// 实例化logger
func InitLogger(logFile string, logErrFile string) (logger *zap.SugaredLogger) {
	encoder := getEncoder()

	writeSyncer := getLogWriter(logFile)
	writeErrSyncer := getLogErrWriter(logErrFile)

	c1 := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	c2 := zapcore.NewCore(encoder, writeErrSyncer, zapcore.ErrorLevel)

	core := zapcore.NewTee(c1, c2)
	return zap.New(core, zap.AddCaller()).Sugar()
}

// 根据文件大小轮转
// 输出到标准输出
func getLogWriter(logFile string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	// 多重输出
	ws := io.MultiWriter(lumberJackLogger, os.Stdout)
	return zapcore.AddSync(ws)
}

// 根据文件大小轮转
// 输出到标准错误
func getLogErrWriter(logErrFile string) zapcore.WriteSyncer {
	lumberJackErrLogger := &lumberjack.Logger{
		Filename:   logErrFile,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	ws := io.MultiWriter(lumberJackErrLogger, os.Stderr)
	return zapcore.AddSync(ws)
}

// 使用json编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// return zapcore.NewConsoleEncoder(encoderConfig)
	return zapcore.NewJSONEncoder(encoderConfig)
}
