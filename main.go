package main

import (
	"bytes"
	"io"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Car  string `json:"car"`
	Num  int    `json:"num"`
}

func main() {
	data := []byte("abcdefghijklmnopq")
	log.Println(len(data))

	r1 := bytes.NewReader(data)

	//  从开头偏移，读取指定长度
	newPos, _ := r1.Seek(2, io.SeekStart)
	log.Println("seek偏移后指针位置:", newPos)

	buf1 := make([]byte, 2)
	n1, _ := r1.Read(buf1)
	log.Println("读取内容与字节数:", string(buf1[:n1]), n1)
	log.Println("Seek操作后r1的内部游标", r1.Size()-int64(r1.Len()))

	r2 := bytes.NewReader(data)
	buf2 := make([]byte, 3)
	n2, _ := r2.ReadAt(buf2, 6)
	log.Println("读取内容与字节数", string(buf2[:n2]), n2)
	log.Println("ReadAt操作后r2的内部游标", r2.Size()-int64(r2.Len()))
}
