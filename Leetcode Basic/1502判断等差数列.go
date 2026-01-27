package main

import (
	"fmt"
)	

func canMakeArithmeticProgression(arr []int) bool {
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-1;j++{
			if arr[j]>arr[j+1]{
			arr[j],arr[j+1]=arr[j+1],arr[j]//冒泡排序
			}
		}
		
	}
	diff:=arr[1]-arr[0]//前两项公差
	for i:=2;i<len(arr);i++{
		if arr[i]-arr[i-1]!=diff{
			return false//后面有任两项差不是公差就不是
		}
	}
	return true
}

//不排序思路：若是等差，(max - min) % (len -1)一定等于0

func main(){
	fmt.Println(canMakeArithmeticProgression([]int{1,2,4}))
}