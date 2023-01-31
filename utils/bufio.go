package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Bufio(path string) []byte {
	var res []byte
	file, err := os.Open(path)
	//er.CheckErr(err)
	if err != nil {
		fmt.Printf("Bufio method -- error:%s ", err.Error())
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)

	for {
		buf := make([]byte, 1024)
		readNum, err := bufReader.Read(buf)
		res = append(res, buf...)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == readNum {
			break
		}
	}
	return res
}
