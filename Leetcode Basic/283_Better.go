package main

import (
	"fmt"
)	
func moveZeroes(nums []int)  {//优化：双指针法
	nonZeroIndex := 0// 用一个指针记录非零元素应该放置的位置
    // 第一步：把所有非零元素移到前面
    for _, num := range nums {
        if num != 0 {
            nums[nonZeroIndex] = num//nums从0开始放数字，把非零数字放进去
            nonZeroIndex++//记录非0数字的个数，指针后移
        }
    }
    // 第二步：把后面的位置全部填充为0
    for i := nonZeroIndex; i < len(nums); i++ {
        nums[i] = 0
    }
}

func main(){
	nums:=[]int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}