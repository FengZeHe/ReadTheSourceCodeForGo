package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// todo 格式化输出
	//var name string
	//var age int
	//log.Println("请输入姓名跟年龄:")
	//
	//scanNum, err := fmt.Scan(&name, &age)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println("获取数量:", scanNum, "姓名:", name, "年龄:", age)

	//var year, month, day int
	//log.Println("请输入日期:")
	//scanf_num, err := fmt.Scanf("%d-%d-%d", &year, &month, &day)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println("获取数量", scanf_num, "年/月/日:", year, month, day)

	//var str string
	//log.Println("请输入一段话:")
	//fmt.Scanln(&str)
	//log.Println("输入:", str)

	//str := "夏天夏天悄悄过去留下 小秘密"
	//var s1, s2 string
	//sscanfNum, err := fmt.Sscanf(str, "%s %s", &s1, &s2)
	//if err != nil {
	//	return
	//}
	//log.Println(sscanfNum, s1, s2)

	wd, _ := os.Getwd()
	fmt.Println("当前工作目录:", wd)

	file, err := os.Open("./fmtdemo/file.txt")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer file.Close()

	var s1, s2 string
	fmt.Fscanf(file, "%s %s", &s1, &s2)
	fmt.Println("从file读取到:", s1, " ", s2)

	// todo  格式化输出

	// todo 错误格式化
}
