package main

import("fmt")

func maxSubArray(nums []int) int {
    current,max1:=nums[0],nums[0]//先计算第0个
	for i:=1;i<len(nums);i++{
		current=max(current+nums[i],nums[i])//看看之前的是小于0还是大于0，小于0就不要，大于0就接上
		max1=max(current,max1)
	}
	return max1
}

func main(){
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}