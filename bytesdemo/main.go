package main

import (
	"bytes"
	"log"
)

func main() {

	// todo 1. 字节切片比较 Equal  Compare
	a, b, c := []byte("abc"), []byte("abc"), []byte("abcd")
	log.Println(bytes.Equal(a, b), bytes.Equal(a, c), bytes.Equal(nil, []byte{}))
	log.Println(bytes.Compare(a, b), bytes.Equal(a, c))

	// todo 2. 是否包含 查找 Contains Index
	str := []byte("abcdefg")
	log.Println(bytes.Contains(str, []byte("abc")))
	log.Println(bytes.Index(str, []byte("f")))

	// todo 3. 前缀/后缀
	log.Println(bytes.HasPrefix(str, []byte("ab")), bytes.HasSuffix(str, []byte("fg")))

	//todo 4. 分割与拼接 Split Join
	str2 := []byte("a,b,c,defg")
	sp := []byte(",")
	log.Println(bytes.Split(str2, sp))

	s3 := [][]byte{}
	s3 = append(s3, []byte("1234"))
	s3 = append(s3, []byte("5678"))
	log.Println(string(bytes.Join(s3, []byte("@"))))

	// todo 5. 修剪 Trim去除首尾
	s5 := []byte("111@qq.com")
	log.Println(string(bytes.Trim(s5, "111")), string(bytes.TrimLeft(s5, "11")), string(bytes.TrimRight(s5, ".com")))

	// todo 6. 大小写转换
	s6 := []byte("abcdefGHI")
	log.Println(string(bytes.ToUpper(s6)), string(bytes.ToLower(s6)))

	// todo 7. 替换 Replace
	s7 := []byte("aaaaabbbccc")
	old, n := []byte("a"), []byte("A")
	log.Println(string(bytes.Replace(s7, old, n, 2)))

	// todo 8. Count
	s8 := []byte("aaaabbbbddd")
	log.Println(bytes.Count(s8, []byte("a")))
}
