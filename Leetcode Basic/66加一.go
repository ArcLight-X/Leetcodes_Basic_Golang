package main

import("fmt"
)

func plusOne(digits []int) []int {
	a:=digits
    for i:=len(a)-1;i>=0;i--{
		a[i]++
        a[i]=a[i]%10// 取余判断是否进位
        if a[i]!=0{// 说明当前位是0-9，没有进位，直接返回结果
            return a
        }//如果有进位，说明上一位也要+1，返回循环
    }
    // 如果循环结束还没返回，说明所有位都进位了
    return append([]int{1}, a...)
	//先创建一个长度为1，值为1的新切片，再原来切片追加到后面，满足如 999 → 1000
}

func main(){
	fmt.Println(plusOne([]int{4,3,9,9}))
}