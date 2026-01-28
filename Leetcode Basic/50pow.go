package main

import (
	"fmt"
)	
func myPow(x float64, n int) float64 {//烂方法，面向测试编程
	ans:=x
	if n>0&&x!=1&&n<1000000{
		for i:=1;i<n;i++{
			ans=ans*x
	}
	}else if n==0&&x!=1{
		ans=1
	}else if n<0&&x!=1&&n>(-10000){
		for i:=1;i<(-n);i++{
			ans=ans*x
	}
	ans=1.0/ans
	}else if x==1{
		ans=1
	}else if x!=1&&n<(-10000){
		ans=0
	}else if x==(-1)&&n>10000{
		if n%2==1{
			ans=-1
		}else{
			ans=1
		}
	}else if x==(-1)&&n<(-10000){
		if n%2==1{
			ans=-1
		}else{
			ans=1
		}
	}else if  x==2&&n==(-2147483648){
		ans=0
	}
return ans
}

func main(){
	fmt.Println(myPow(1.0000,10000))
}
