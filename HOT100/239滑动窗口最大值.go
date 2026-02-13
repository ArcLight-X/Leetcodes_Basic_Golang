package main

import (
	"container/list"
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums)==0||k<= 0{//边界条件处理
		return []int{}
	}
	if k==1 {
		return nums//此时最大值就是元素本身
	}

	//用list实现双端队列
	deque:=list.New()
	result:=make([]int,0,len(nums)-k+1)

	for i, num := range nums {
		//移除超出窗口范围的索引（窗口左边界：i-k）
		for deque.Len()>0 &&deque.Front().Value.(int)<=i-k{
			deque.Remove(deque.Front())
		}

		//保证队列单调性：移除尾部所有值小于当前元素的索引
		for deque.Len()>0&&nums[deque.Back().Value.(int)]<num{
			deque.Remove(deque.Back())
		}

		//将当前索引加入队列尾部
		deque.PushBack(i)

		//窗口形成（i >= k-1）时，取队首为最大值
		if i>=k-1 {
			result=append(result, nums[deque.Front().Value.(int)])
		}
	}

	return result
}



func main(){
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3))
}
