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
	//合并有序,同21

	return mergeTwoLists(left, right)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
	/*1. dummy := &ListNode{} → 空虚拟头（Next=nil）
	用途：新建一条全新的链表
	手里没有现成链表，要从零拼接
	dummy := &ListNode{Next: head} → 挂原头的虚拟头
	用途：改造已有的旧链表*/
    cur := dummy
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            cur.Next = list1
            list1 = list1.Next
        } else {
            cur.Next = list2
            list2 = list2.Next
        }
        cur = cur.Next
    }
    if list1 != nil {
        cur.Next = list1
    } else {
        cur.Next = list2
    }
    return dummy.Next
}