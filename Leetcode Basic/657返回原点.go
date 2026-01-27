package main

import (
	"fmt"
)	

func judgeCircle(moves string) bool {
	var i,r,l,u,d int
    for i=0;i<len(moves);i++{
		if moves[i]==82{//通过ASCII码识别字符
			r+=1
		}else if moves[i]==76{
			l+=1
		}else if moves[i]==85{
			u+=1
		}else if moves[i]==68{
			d+=1
		}
	}
	if r==l&&u==d{//能回到原点，则上（的步数）=下，左=右
		return true
	}else{
		return false
	}
}

func main(){
	fmt.Println(judgeCircle("LL"))
}