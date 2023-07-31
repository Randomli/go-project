package main

import (
	"henry.com/go-project/leetcode_array"
	"henry.com/go-project/logger"
)

func main() {
	// ref.Run()
	// ref.Run2()
	// ref.Run3()
	// ref.Run4()
	// gor.Run()
	// gor.Run1()
	// gor.Run2()
	// gor.Run3()
	// gor.Run5()
	// FetchUrlDemo2()
	// libstudy.TFprint()
	// libstudy.TSprint()
	// libstudy.TErrorf()
	// libstudy.TPlaceholder()
	// libstudy.TTimeDemo()
	// libstudy.TTimezoneDemo()
	// libstudy.TUnixTimeDemo()
	// libstudy.TPrintTime()
	// libstudy.TUseMicrosecond()
	//
	// defer logger.Sync()
	// logger.Info("info级别日志")
	// logger.Error("error级别日志")
	// logger.Errorw("运行故障",
	// 	"reason", "不知道",
	// 	"new_key", "1",
	// )

	// code := 0
	// defer func() {
	// 	os.Exit(code)
	// }()
	// subCmd := os.Args[1]
	// switch subCmd {
	// case "server":
	// 	fmt.Println("run server")
	// 	RunServer()
	// case "client":
	// 	fmt.Println("run client")
	// 	RunClient()
	// case "stick_server":
	// 	fmt.Println("run stick_server")
	// 	RunStickServer()
	// case "stick_client":
	// 	fmt.Println("run stick_client")
	// 	RunStickClient()
	// default:
	// 	fmt.Println("子命令不支持")
	// 	code = 1
	// }

	logger := logger.InitLogger(LOGFILE, LOGERRFILE)
	defer logger.Sync()
	logger.Info("info级别日志")
	leetcode_array.Run()
}
