package main

import("fmt"
"strconv"
)
//此方法无法处理大数，会溢出
func multiply(num1 string, num2 string) string {
    x1,x2:=0,0
	m,n:=1,1
	for i:=len(num1)-1;i>=0;i--{
		x1=x1+int(num1[i]-'0')*m
		m=m*10
	}
	for j:=len(num2)-1;j>=0;j--{
		x2=x2+int(num2[j]-'0')*n
		n=n*10
	}
	ans:=x1*x2
	return strconv.Itoa(ans)

}

func main(){
	fmt.Println(multiply("123","456"))
}