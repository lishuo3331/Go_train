package main

import (
	"fmt"
	"os"
	)

func main()  {
	var s, sep string
	//读入命令行并输出 以空格为元素分界
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
