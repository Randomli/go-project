package ref

import (
	"fmt"
	"reflect"
)

func reflectType1(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}
func Run() {
	var a float32 = 3.14
	reflectType1(a) // type:float32
	var b int64 = 100
	reflectType1(b) // type:int64
}

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func Run2() {
	var a *float32
	var b myInt
	var c rune
	reflectType(a)
	reflectType(b)
	reflectType(c)

	type person struct {
		name string
		age  int
	}
	type book struct {
		title string
	}

	var d = person{
		name: "h",
		age:  18,
	}
	var e = book{title: "golang"}

	reflectType(d)
	reflectType(e)

}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
func Run3() {
	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}

func Run4() {
	var a *int

	bool2string := map[bool]string{
		true:  "存在",
		false: "不存在",
	}

	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())

	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())

	b := struct{}{}

	fmt.Printf("%v结构体成员\n", bool2string[reflect.ValueOf(b).FieldByName("abc").IsValid()])
	fmt.Printf("%v结构体方法\n", bool2string[reflect.ValueOf(b).MethodByName("abc").IsValid()])

	c := map[string]int{}

	fmt.Printf("map中%v'娜扎'键\n", bool2string[reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid()])
}
