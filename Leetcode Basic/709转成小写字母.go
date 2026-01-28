package main

import "fmt"

func toLowerCase(s string) string {
	bytes := []byte(s)//转为可变字节数组
    for i:=0;i<len(s);i++{
		if bytes[i]>='A'&&bytes[i]<='Z'{//>=65 <=90也可以
			bytes[i]+=32
		}
	}
	return string(bytes)
}

func main(){
	fmt.Println(toLowerCase("LOVELY"))
}