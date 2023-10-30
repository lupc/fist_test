package test

import (
	"fmt"
	"io"
	"time"

	"github.com/lupc/go-rollingwriter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger2() (l *zap.Logger, err error) {
	var logLevel = zapcore.DebugLevel
	// err = logLevel.UnmarshalText([]byte("debug"))

	warnIoWriter := getWriter2()
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
func getWriter2() io.Writer {

	var cfg = rollingwriter.Config{
		LogPath:                "./logs2/{yyyy-MM}",
		TimeTagFormat:          "20060102150405",
		FileName:               "log_{HH}",
		FileExtension:          "log",
		MaxRemain:              0,             // disable auto delete
		RollingPolicy:          2,             // TimeRotate by default
		RollingTimePattern:     "0 0 0 * * *", // Rolling at 00:00 AM everyday
		RollingVolumeSize:      "10M",
		WriterMode:             "async",
		BufferWriterThershould: 64,
		Compress:               false,
	}
	var w, err = rollingwriter.NewWriterFromConfig(&cfg)

	if err != nil {
		panic(err)
	}
	return w
}

func RunZapLogRote2() {

	log, err := GetLogger2()
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

func RunZapLogRote2Benchmark() {

	log, err := GetLogger2()

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
