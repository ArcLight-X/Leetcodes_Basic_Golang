package main

import("fmt")

func arraySign(nums []int) int {
	posi:=0;//直接算有可能超范围，统计正负数
	nega:=0
    for i:=0;i<len(nums);i++{
		if nums[i]>0{
			posi++
		}else if nums[i]<0{
			nega++
		}else if nums[i]==0{
			return 0
		}
	}
	if nega%2==1{
		return -1
	}else if nega%2==0{
		return 1
	}
	return 2
}

func main(){
	fmt.Println(arraySign([]int{-1,-2,-3,-4,3,2,1}))
}