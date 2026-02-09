package main

import (
	"fmt"
)	
/*
func moveZeroes(nums []int)  {
    for i:=0;i<len(nums)-1;i++{
		for j:=0;j<len(nums)-1;j++{
			if nums[j]==0&&nums[j+1]!=0{//如果i个是0.i+1个不是0，就交换
			nums[j],nums[j+1]=nums[j+1],nums[j]
		}
		}
	}
}*///类冒泡排序，效率低

//双指针法：
func moveZeroes(nums []int)  {
	slow:=0//慢针，记录非0元素位置
	for fast:=0;fast<len(nums);fast++{//快针，遍历数组找非0元素
		if nums[fast]!=0{//找到非0元素
			nums[slow],nums[fast]=nums[fast],nums[slow]//快慢针指向的元素交换
			slow++//为下一个非0元素腾位置
		}
	}
}

func main(){
	nums:=[]int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}
/*初始状态: [0,1,0,3,12] slow=0, fast=0
nums[0]=0 → 跳过，slow=0, fast=1

nums[1]=1 ≠ 0 → 交换 nums[0]和nums[1]: [1,0,0,3,12], slow=1, fast=2
nums[2]=0 → 跳过，slow=1, fast=3

nums[3]=3 ≠ 0 → 交换 nums[1]和nums[3]: [1,3,0,0,12], slow=2, fast=4
nums[4]=12 ≠ 0 → 交换 nums[2]和nums[4]: [1,3,12,0,0], slow=3, fast=5

结束：[1,3,12,0,0]*/