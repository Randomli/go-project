module henry.com/go-project

go 1.18

replace henry.com/go-project/ref => ./ref

replace henry.com/go-project/gor => ./gor

// 本地引入
replace libstudy => ./lib_study

require (
	go.uber.org/zap v1.24.0
	golang.org/x/sync v0.3.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	libstudy v0.0.0-00010101000000-000000000000
)

require (
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
)
