package test

import (
	"fmt"
	"io"
	"time"

	rotatelogs "github.com/lupc/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger() (l *zap.Logger, err error) {
	var logLevel = zapcore.DebugLevel
	// err = logLevel.UnmarshalText([]byte("debug"))

	warnIoWriter := getWriter("./logs/%Y-%m/test_%Y%m%d.log")
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
func getWriter(logFile string) io.Writer {
	// 保存30天内的日志，每24小时(整点)分割一次日志
	writer, err := rotatelogs.New(
		// logFile+".%Y%m%d",
		logFile, //每天
		rotatelogs.WithLinkName("./logs/current.log"), //生成软链，指向最新日志文件
		rotatelogs.WithRotationTime(24*time.Hour),     //最小为1分钟轮询。默认60s  低于1分钟就按1分钟来
		rotatelogs.WithRotationCount(0),               //设置3份 大于3份 或到了清理时间 开始清理
		// rotatelogs.WithMaxAge(1*time.Minute),
		rotatelogs.WithRotationSize(10*1024*1024), //设置100MB大小,当大于这个容量时，创建新的日志文件

	)

	if err != nil {
		panic(err)
	}
	return writer
}

func RunZapLogRote() {

	log, err := GetLogger()
	if err == nil {
		for {

			log.Debug("Debug日志内容。。", zap.String("key1", "string"), zap.Float32("key2", 223.444))
			log.Info("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			log.Warn("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			log.Error("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			// log.Fatal("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

			// time.Sleep(5 * time.Second)
		}
	}
}

func RunZapLogRoteBenchmark() {

	log, err := GetLogger()

	if err == nil {

		var t1 = time.Now()
		for i := 0; i < 10*10000; i++ {
			log.Debug("Debug日志内容。。", zap.String("key1", "string"), zap.Float32("key2", 223.444))
			log.Info("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			log.Warn("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			log.Error("11111111111111111aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		}
		var ts = time.Since(t1)
		fmt.Printf("耗时 ts: %v\n", ts)
	}

}
