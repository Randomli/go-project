package main

// ref "henry.com/go-project/ref"

// import (
// 	"flag"
// 	"fmt"
// )
import (
	"fmt"
	"os"

	"libstudy"
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
	logger.Info("test")
	libstudy.TUseMicrosecond()
	code := 0
	defer func() {
		os.Exit(code)
	}()
	subCmd := os.Args[1]
	switch subCmd {
	case "server":
		fmt.Println("run server")
		RunServer()
	case "client":
		fmt.Println("run client")
		RunClient()
	case "stick_server":
		fmt.Println("run stick_server")
		RunStickServer()
	case "stick_client":
		fmt.Println("run stick_client")
		RunStickClient()
	default:
		fmt.Println("子命令不支持")
		code = 1
	}

}
