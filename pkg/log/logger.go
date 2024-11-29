package log

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once     sync.Once          // 用于确保单例初始化的工具
	instance *zap.SugaredLogger // 单例实例
)

// getLoggerInstance 是一个内部方法，用于创建 Logger
func getLoggerInstance() *zap.SugaredLogger {
	// 自定义日志配置
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel), // 设置日志级别
		Development:      false,                                // 是否是开发模式
		Encoding:         "json",                               // 输出格式：json 或 console
		OutputPaths:      []string{"stdout"},                   // 输出位置
		ErrorOutputPaths: []string{"stderr"},                   // 错误输出位置
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			MessageKey:    "msg",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			EncodeTime:    zapcore.ISO8601TimeEncoder,  // 时间格式
			EncodeLevel:   zapcore.CapitalLevelEncoder, // 日志级别大写
			EncodeCaller:  zapcore.ShortCallerEncoder,  // 文件名:行号
		},
	}

	// 创建 Logger 实例
	logger, err := config.Build()
	if err != nil {
		panic("failed to create logger: " + err.Error())
	}

	return logger.Sugar()
}

// GetLogger 提供全局访问单例 Logger 的方法
func GetLogger() *zap.SugaredLogger {
	once.Do(func() {
		instance = getLoggerInstance()
	})
	return instance
}
