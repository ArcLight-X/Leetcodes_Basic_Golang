package main

import("fmt"
		"strconv"
)

func calPoints(operations []string) int {
	sum:=0
	a:=[]int{}
	for _,v :=range operations{
		i:=len(a)//实时确定长度
		switch v{
		case "+":
			a=append(a,a[i-1]+a[i-2])
		case "C":
			a=a[:i-1]//把最后一个删了
		case "D":
			a=append(a,2*a[i-1])
		default:
			x,_:=strconv.Atoi(v)//第二个返回值是转换失败的报错
			a=append(a,x)
		}
	}
	for j:=0;j<len(a);j++{
		sum+=a[j]
	}
	return sum
}

func main(){
	fmt.Println(calPoints([]string{"5","2","C","D","+"}))
}