package libstudy

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func TFprint() {
	fmt.Fprintln(os.Stdout, "向标准输出写入内容", "\t内容2", "\t内容3")
	// 没有文件时新建 只写打开 截断文件 0644权限
	f, err := os.OpenFile("./abc.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("打开文件出错, err:", err)
		return
	}
	name := "henry"

	fmt.Fprintln(f, "往文件中写入信息:", name)
	fmt.Fprintln(f, "往文件中写入信息:", name)
	fmt.Fprintln(f, "往文件中写入信息:", name)
	fmt.Fprintln(f, "往文件中写入信息:", name)
	fmt.Fprintf(f, "往文件中写入信息:%s", name)
	fmt.Fprintf(f, "往文件中写入信息:%s", name)
	fmt.Fprintf(f, "往文件中写入信息:%s", name)

}

func TSprint() {
	s1 := fmt.Sprint("henry")

	name := "henry"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)

	s3 := fmt.Sprintln("henry")

	fmt.Println(s1, s2, s3)
}

func TErrorf() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e)
	fmt.Println(w)
}

func TPlaceholder() {
	fmt.Printf("unicode %v\n", "ll")
}

func TTimeDemo() {
	now := time.Now()
	fmt.Printf("current time: %v\n", now)

	year := now.Year()
	month_str := fmt.Sprintf("%02d", int(now.Month()))
	day := now.Day()

	hour := now.Hour()
	minite := now.Minute()
	second := now.Second()
	fmt.Println(year, month_str, day, hour, minite, second)
}

var FixedZone *time.Location = time.FixedZone("Beijing Time", int((8 * time.Hour).Seconds()))

// timezoneDemo 时区示例
func TTimezoneDemo() {
	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区
	// 例如，加载纽约所在的时区
	newYork, err := time.LoadLocation("America/New_York") // UTC-05:00
	if err != nil {
		fmt.Println("load America/New_York location failed", err)
		return
	}
	fmt.Println()
	// 加载上海所在的时区
	//shanghai, err := time.LoadLocation("Asia/Shanghai") // UTC+08:00
	// 加载东京所在的时区
	//tokyo, err := time.LoadLocation("Asia/Tokyo") // UTC+09:00

	// 创建时间对象需要指定位置。常用的位置是 time.Local（当地时间） 和 time.UTC（UTC时间）。
	//timeInLocal := time.Date(2009, 1, 1, 20, 0, 0, 0, time.Local)  // 系统本地时间
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	sameTimeInNewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newYork)

	// 北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)

	// 纽约（西五区）比UTC晚5小时，所以上面两个时间看似差了5小时，但表示的是同一个时间
	timesAreEqual = timeInUTC.Equal(sameTimeInNewYork)
	fmt.Println(timesAreEqual)
}

func TUnixTimeDemo() {
	// timestampDemo 时间戳
	// now := time.Now() // 获取当前时间
	now := time.Now().In(FixedZone)
	fmt.Printf("current time: %v\n", now)
	timestamp := now.Unix()  // 秒级时间戳
	milli := now.UnixMilli() // 毫秒时间戳 Go1.17+
	micro := now.UnixMicro() // 微秒时间戳 Go1.17+
	nano := now.UnixNano()   // 纳秒时间戳
	fmt.Println(timestamp, milli, micro, nano)
}

func TPrintTime() {
	// 获取当前时间，格式化输出为2017/06/19 20:30:05格式。
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))
}

func sleepDemo() {
	time.Sleep(time.Microsecond * 1)

}
func TUseMicrosecond() {
	// 编写程序统计一段代码的执行耗时时间，单位精确到微秒。
	start := time.Now().UnixNano()
	sleepDemo()
	end := time.Now().UnixNano()
	fmt.Printf("cost: %dus\n", (end-start)/1000)

}
