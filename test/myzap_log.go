package test

import (
	"fmt"
	"time"

	"github.com/lupc/go-myzap"
	"go.uber.org/zap"
)

func GetLoggerMyZap() (l *zap.Logger, err error) {
	var c = myzap.NewConfigByName("myzap")
	c.IsLogToConsole = false
	l = c.BuildLogger()
	return l, nil
}

func RunMyZap() {

	log, err := GetLoggerMyZap()
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

func RunMyZapBenchmark() {

	log, err := GetLoggerMyZap()

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
