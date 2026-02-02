package main

import("fmt")

func largestPerimeter(nums []int) int {
    for i:=0;i<len(nums)-1;i++{
		for j:=0;j<len(nums)-1;j++{
			if nums[j]>nums[j+1]{
				nums[j],nums[j+1]=nums[j+1],nums[j]
			}
		}
	}
	for i:=len(nums)-1;i>=2;i--{//从最大元素倒序检查
		if nums[i-2]+nums[i-1]>nums[i]{//数组已经升序了，满足下面条件就能组成三角形
			return nums[i-2]+nums[i-1]+nums[i]
		}
	}
	return 0
		
}

func main(){
	fmt.Println(largestPerimeter([]int{1,2,2}))
}