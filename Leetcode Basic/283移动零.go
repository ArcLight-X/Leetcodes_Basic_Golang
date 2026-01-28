package main

import (
	"fmt"
)	

func moveZeroes(nums []int)  {
    for i:=0;i<len(nums)-1;i++{
		for j:=0;j<len(nums)-1;j++{
			if nums[j]==0&&nums[j+1]!=0{//如果i个是0.i+1个不是0，就交换
			nums[j],nums[j+1]=nums[j+1],nums[j]
		}
		}
	}
}
func main(){
	nums:=[]int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}