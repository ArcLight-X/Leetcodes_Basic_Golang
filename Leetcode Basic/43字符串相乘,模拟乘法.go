package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1=="0"||num2=="0"{
		return "0"
	}
    ans:=make([]int,len(num1)+len(num2))//把结果存在数组中
	digit1,digit2,sum,mul:=0,0,0,0
	for i:=len(num1)-1;i>=0;i--{
		digit1=int(num1[i]-'0')
		for j:=len(num2)-1;j>=0;j--{
			digit2=int(num2[j]-'0')
			mul=digit1*digit2
			//本位落在 ans[i+j+1]，进位落在 ans[i+j]
			sum=mul+ans[i+j+1]//模拟乘法竖式中每一次乘法计算的结果
			ans[i+j+1]=sum%10//取计算结果的个位，存到本位上
			ans[i+j]=ans[i+j]+sum/10//进位,存到进位上
		}
	}
	result := ""
	for _, val:=range ans {
		// 跳过开头的连续0
		if len(result)==0 && val==0 {
			continue
		}
		result+=strconv.Itoa(val)
	}

	return result
}

func main(){
	fmt.Println(multiply("123","456"))
}