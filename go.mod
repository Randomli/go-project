module henry.com/go-project

go 1.18

replace henry.com/go-project/ref => ./ref

replace henry.com/go-project/gor => ./gor

// 本地引入
replace henry.com/go-project/leetcode_array => ./leetcode/array

replace henry.com/go-project/logger => ./logger

replace henry.com/go-project/libstudy => ./lib_study

require (
	golang.org/x/sync v0.3.0
	henry.com/go-project/leetcode_array v0.0.0-00010101000000-000000000000
	henry.com/go-project/logger v0.0.0-00010101000000-000000000000
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)
