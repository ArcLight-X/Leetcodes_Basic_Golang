package main

import("fmt")
//对于任意位置i，能接到的雨水量 = min(左边最高,右边最高)-height[i]
//双指针L: 使用left和right两个指针从两端向中间移动，同时维护leftMax和rightMax
func trap(height []int) int {
    left,right:=0,len(height)-1
	leftMax,rightMax:=0,0
	sum:=0
	for left<right{
		leftMax=max(height[left],leftMax)
		rightMax=max(height[right],rightMax)
		if leftMax<=rightMax{
		//当 leftMax<=rightMax时,虽然不知道右侧更远处是否有更高的柱子，但可以确定当前位置的积水高度受到左侧最大值限制
		//因为积水高度取决于 min(左侧最大值, 右侧最大值) ，而左侧最大值是确定的，且≤右侧最大值
			sum=sum+leftMax-height[left]//积水量则是左侧柱子最大值-当前指针指向柱子高度
			left++
		}else{//右侧同理
			sum=sum+rightMax-height[right]
			right--
		}
	}
	return sum
}

func main(){
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}
