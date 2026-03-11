package main

import("fmt"

)

type ListNode struct {
    Val int
    Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	//空链表或单节点，是回文
    if head==nil||head.Next==nil{
		return true
	}
	//先用快慢针找中点
	slow:=head
	fast:=head
	for fast.Next!=nil&&fast.Next.Next!=nil{//要提前停止
		slow=slow.Next
		fast=fast.Next.Next
		//慢走一步，快走两步，结束时慢针一定在中点或中点左侧
	}
	//把后半段反转，用头插法
	var prev *ListNode //头插法的头，初始为nil
    curr:=slow.Next
	for curr!=nil{
		temp:=curr.Next
		curr.Next=prev
		prev=curr
		curr=temp
	}
	//前半段和反转的后半段比较
	p1:=head
	p2:=prev
	for p1!=nil{//p2是后半段，更短
		if p1.Val!=p2.Val{
			return false
		}
		p1=p1.Next
		p2=p2.Next
	}
	return true

}
