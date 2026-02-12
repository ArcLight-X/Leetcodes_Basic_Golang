package main

import("fmt"

)
//三次翻转

func reverse(a []int,x int,y int)[]int{//辅助函数：翻转第x到y个
		left,right:=x,y
		for left<right{
			a[left],a[right]=a[right],a[left]
        	left++
        	right--
		}
	return a
}
func rotate(nums []int, k int) []int {
    
	k=k%len(nums)//翻超过一整圈就没意义
	nums=reverse(nums,0,len(nums)-1)//第一次：整体翻转
	nums=reverse(nums,0,k-1)//        第二次：前k个翻转
	nums=reverse(nums,k,len(nums)-1)//第三次：后n-k个翻转
	return nums
}

func main(){
	fmt.Println(rotate([]int{1,2,3,4,5,6,7},3))
}