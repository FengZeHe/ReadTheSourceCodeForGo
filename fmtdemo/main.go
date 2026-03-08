package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	CodeOne = errors.New("code1")
	CodeTwo = errors.New("code2")
)

func main() {
	// 格式化输出
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

	//wd, _ := os.Getwd()
	//fmt.Println("当前工作目录:", wd)
	//
	//file, err := os.Open("./fmtdemo/file.txt")
	//if err != nil {
	//	log.Println(err)
	//	panic(err)
	//}
	//defer file.Close()
	//
	//var s1, s2 string
	//fmt.Fscanf(file, "%s %s", &s1, &s2)
	//fmt.Println("从file读取到:", s1, " ", s2)

	// 格式化输出
	//name := "dawei"
	//fmt.Printf("hello %v \n", name)
	//fmt.Printf("%+v \n", name)
	//fmt.Printf("%#v \n", name)
	//fmt.Printf("%T \n", name)
	//
	//fmt.Printf("%t \n", true)
	//fmt.Printf("保留2位小数: %.2f, 科学计数法: %e\n", 1.2345, 100000.0)

	// 错误格式化,生成简单的错误
	//code := 101
	//err := fmt.Errorf("error: %d", code)
	//fmt.Println(err)
	//
	//err2 := AddMsg()
	//fmt.Println(err2)

	err := addMsg()
	if err != nil {
		fmt.Println(err)

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Error", err.Error())
		}

		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			fmt.Printf("操作:%s, 路径:%s, 底层错误=%v\n", pathErr.Op, pathErr.Path, pathErr.Err)
		}
	}

}

func addMsg() error {
	_, err := os.Open("abc.txt")
	if err != nil {
		return fmt.Errorf("open abc.txt: %w", err)
	}
	return nil
}
