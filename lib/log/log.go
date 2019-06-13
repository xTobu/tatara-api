package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// Logger : exported for convenience
	Logger *zap.Logger
)

// Init : Initialed
func Init() {

	Logger = initLogger("error")

}

// GetLogger : Get Zap logger
func GetLogger() *zap.Logger {

	return Logger

}

func initLogger(logLevel string) *zap.Logger {

	var (
		logger *zap.Logger
	)

	hook := lumberjack.Logger{
		Filename:   "./logs/tatara-api.log", // 日誌文件路徑
		MaxSize:    128,                     // 每個日誌最大容量：M
		MaxBackups: 30,                      // 最多保存多少個備份
		MaxAge:     7,                       // 最多保存天數
		Compress:   true,                    // 是否壓縮
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 大寫寫編碼器
		EncodeTime:     zapcore.ISO8601TimeEncoder,  // ISO8601 UTC 時間格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 簡短路徑的編碼器
		EncodeName:     zapcore.FullNameEncoder,
	}

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "info":
	default:
		level = zap.InfoLevel
	}

	// 設置日誌級別
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 編碼器配置
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stderr), // for console
			zapcore.AddSync(&hook),     // for lumberjack.logger
		),
		atomicLevel, // 日誌級別
	)

	// 開啟開發模式，堆棧跟踪
	caller := zap.AddCaller()

	// 開啟文件及行號
	development := zap.Development()

	// 設置初始化字段
	// filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// Logger := zap.New(core, caller, development, filed)

	// 建置日誌
	logger = zap.New(core, caller, development)

	/* Demo for using */
	// logger.Info("無法獲取網址",
	// 	zap.String("url", "http://www.baidu.com"),
	// 	zap.Int("attempt", 3),
	// 	zap.Duration("backoff", time.Second))

	return logger
}
