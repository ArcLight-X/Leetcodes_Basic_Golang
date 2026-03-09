package main

import("fmt"

)

type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}
	dummy:=&ListNode{Next:head}//创建一个虚拟头节点
	curr := dummy
	for curr.Next!=nil&&curr.Next.Next!=nil {//确保两个节点可以交换
		//保存待交换的两个节点
		first:=curr.Next
		second:=curr.Next.Next

		first.Next = second.Next //第一个节点指向第二个节点的下一个节点
		second.Next = first      //第二个节点指向第一个节点
		curr.Next = second       //前驱节点指向交换后的新头节点

		curr=first//移动当前指针，准备下一组交换
	}
	return dummy.Next
}