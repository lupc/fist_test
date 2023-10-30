package test

import (
	"fmt"
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func GetLogger3() (l *zap.Logger, err error) {
	var logLevel = zapcore.DebugLevel
	// err = logLevel.UnmarshalText([]byte("debug"))

	warnIoWriter := getWriter3()
	// _ = os.Mkdir("./logs", 0755)
	// var writer = getLogWriter()
	var encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	var coreFile = zapcore.NewCore(encoder, zapcore.AddSync(warnIoWriter), logLevel)
	// var coreConsole = zapcore.NewCore(encoder, os.Stdout, logLevel)

	// var core = zapcore.NewTee(coreFile, coreConsole)
	var core = zapcore.NewTee(coreFile)
	zaplog := zap.New(core, zap.AddStacktrace(zap.ErrorLevel), zap.AddCaller())
	return zaplog, nil
}

// 日志文件切割
func getWriter3() io.Writer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs3/test.log", // 文件位置
		MaxSize:    1,                  // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     30,                 // 保留旧文件的最大天数
		MaxBackups: 0,                  // 保留旧文件的最大个数
		Compress:   false,              // 是否压缩/归档旧文件
	}
	return lumberJackLogger
}

func RunZapLogRote3() {

	log, err := GetLogger3()
	if err == nil {
		for {

			log.Debug("Debug日志内容。。", zap.String("key1", "string"), zap.Float32("key2", 223.444))
			log.Info("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			log.Warn("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			log.Error("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			// log.Fatal("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		}

	}
}

func RunZapLogRote3Benchmark() {

	log, err := GetLogger3()

	if err == nil {

		var t1 = time.Now()
		for i := 0; i < 10*10000; i++ {
			log.Debug("Debug日志内容。。", zap.String("key1", "string"), zap.Float32("key2", 223.444))
			log.Info("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			log.Warn("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			log.Error("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			// time.Sleep(1 * time.Millisecond)
		}
		var ts = time.Since(t1)
		fmt.Printf("耗时 ts: %v\n", ts)

	}

}
