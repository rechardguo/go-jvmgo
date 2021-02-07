package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	mapDemo()
	errorDemo()
	强转类型()
	arrayDemo()

	var i1 int = 1
	var i2 float32 = 1
	println(i1)
	println(i2)
	fmt.Printf("%v \n", math.Float32bits(i2))
	my := &User{}
	my.setName("rechard")
	println(my.name)

	println("hello,go")
	u := newUser("rechard", 35)
	println(u.String())

	i32 := 0x234234
	println(int(i32))

	fileName := "path"
	_, error := os.Stat(fileName)
	if error != nil {
		panic("file not found")
	}

}

func arrayDemo() {
	arr := []int{1, 4, 5, 2, 3}
	arr = arr[0:5] //[begin,end] begin包含， end不包含
	// _表示不使用,如果写成index，则表示使用
	for _, v := range arr {
		println(v)
	}
	//i = index
	//v = value
	for i, _ := range arr {
		println(i)
	}

	arr2 := make([]int, 5)
	arr2[0] = 12

	traverse(arr)
	slice(arr)
}

func 强转类型() {
	animal := newAnmial("dog")
	//取得实际类型，类似java里的 dog.getClass()
	switch animal.(type) {
	case *Dog:
		print(animal.Growl())
		//直接这样是不行的
		//dog.Eat()
		//强转成为Dog
		dog := animal.(*Dog)
		println(dog.Eat())
	}
}

func errorDemo() {
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}
func mapDemo() {
	//创建map的方式1
	var countryCapitalMap map[string]string = make(map[string]string)
	countryCapitalMap["china"] = "beijing"
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	capital, exist := countryCapitalMap["America"]
	if exist {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
	//创建map的方式2
	//直接创建
	idPersonMap := map[int]string{1: "rechard", 2: "tom", 3: "james"}
	for id := range idPersonMap {
		fmt.Println(id, "=", idPersonMap[id])
	}
}

func catchErr() {
	if r := recover(); r != nil {
		err := r.(error)
		fmt.Printf("error:%s", err.Error())
	}

}
