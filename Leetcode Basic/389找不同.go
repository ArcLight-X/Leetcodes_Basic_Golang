package main
import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
	var res byte
    for _, c := range s {//异或运算
        res ^= byte(c)//相同的字符保存，相异的抵消
    }
    for _, c := range t {
        res ^= byte(c)//与上面保存的进行异或，都会抵消，留下了的就是多的
    }
    return res
}

func main(){
	fmt.Printf("%c",findTheDifference("abcd","abcde"))
}