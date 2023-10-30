module fist_test

go 1.20

//replace github.com/arthurkiller/rollingwriter => github.com/lupc/go-rollingwriter
replace github.com/lupc/go-myzap => github.com/lupc/go-myzap v1.0.2

require (
	github.com/lupc/go-file-rotatelogs v1.0.4-0.20231019101313-0a16c6086d6e
	github.com/lupc/go-myzap v1.0.3
	github.com/lupc/go-rollingwriter v1.1.0
)

require (
	github.com/robfig/cron v1.2.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
)

require (
	github.com/connerdouglass/go-geo v1.0.1
	github.com/golang/geo v0.0.0-20230421003525-6adc56603217
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/zerolog v1.31.0
	github.com/twpayne/go-geom v1.5.2
	go.uber.org/zap v1.26.0
	golang.org/x/sys v0.13.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)
