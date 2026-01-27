package main
import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	str1 := word1
	str2 := word2
	a:=len(str1)
	b:=len(str2)
	var c int
	if a>=b{//找出c大小,c前面的字符交替合并，c后面的加到最后
		c=b
	}else{
		c=a
	}
	var str3 string
	for id:=0;id<c;id++{//交替合并
		str3+=string(str1[id])
		str3+=string(str2[id])
	}
	str3 += str1[c:] //加上c后面的，没有的就加“”，不影响
	str3 += str2[c:]
	return str3
}

func main(){
	fmt.Println(mergeAlternately("ab","pqrs"))
}
