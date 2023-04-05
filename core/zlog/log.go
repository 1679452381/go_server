/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package zlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go_server/global"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"runtime"
)

var logger *zap.Logger

func getFileLogWriter() (syncer zapcore.WriteSyncer) {
	// 使用 lumberjack 实现 日志切分
	lumberJackLogger := &lumberjack.Logger{
		Filename:   global.Config.Logger.Director, //文件路径
		MaxSize:    100,                           // 单个文件最大100M
		MaxBackups: 60,                            // 多于 60 个日志文件后，清理较旧的日志
		MaxAge:     1,                             // 一天一切割
		Compress:   false,                         // 决定了回滚的文件是否需要压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}
func Init() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	//设置日志记录中时间的格式 "2006-01-02T15:04:05.000Z0700"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志Encoder 还是JSONEncoder，把日志行格式化成JSON格式的
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	//file, _ := os.OpenFile("./zap_study/zap_log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 644)
	fileWriteSyncer := getFileLogWriter()

	level, err := zapcore.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		panic(fmt.Errorf("get yamlConf err:%s", err))
	}
	core := zapcore.NewTee(
		// 同时向控制台和文件写日志， 生产环境记得把控制台写入去掉，日志记录的基本是Debug 及以上，生产环境记得改成Info
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level),
		zapcore.NewCore(encoder, fileWriteSyncer, level),
	)
	logger = zap.New(core)
	return logger
}

func Info(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Info(message, fields...)
}
func Debug(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Debug(message, fields...)
}
func Error(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Error(message, fields...)
}
func Warn(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Warn(message, fields...)
}

func getCallerInfoForLog() (callerFields []zap.Field) {
	pc, file, line, ok := runtime.Caller(2) // 回溯两层，拿到写日志的调用方的函数信息
	if !ok {
		return
	}
	funName := runtime.FuncForPC(pc).Name()
	funName = path.Base(funName) //Base函数返回路径的最后一个元素，只保留函数名

	callerFields = append(callerFields,
		zap.String("func", funName),
		zap.String("file", file),
		zap.Int("line", line))
	return
}
