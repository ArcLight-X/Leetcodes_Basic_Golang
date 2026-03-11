package main

import("fmt"

)

type ListNode struct {
    Val int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
    if head==nil||head.Next==nil{
		return head
	}
	slow, fast:=head, head.Next//快慢针找中点
    for fast!=nil&&fast.Next!=nil {
        slow=slow.Next
        fast=fast.Next.Next
    }
    mid :=slow.Next
    slow.Next =nil//切断链表，变成head和mid为头节点的两段
	//归并排序，使其有序
	left:=sortList(head)
	right:=sortList(mid)
    //迭代法合并有序链表
    dummy := &ListNode{}
    curr := dummy
    for left != nil && right != nil {//比较两个链表的【当前第一个节点】，谁小就先拼谁
        if left.Val < right.Val {
            curr.Next = left// 拼left的当前节点：把curr的下一个指向left
            left = left.Next// left的第一个已经拼完了，往后走一步，处理下一个节点
        } else {
            curr.Next = right
            right = right.Next
        }
        curr = curr.Next//结果的第一个拼完了，往后走
    }
    //处理剩余节点（一个链表走完了，另一个还剩节点，直接全部拼上去）
    if left != nil {
        curr.Next = left
    }
    if right != nil {
        curr.Next = right
    }
    return dummy.Next
}