package main

import("fmt"
)
//双指针不能用，不是数组
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode //头插法的头，初始为nil
    curr := head       //当前节点，初始指向头节点

    for curr != nil {
        temp := curr.Next //保存当前节点的下一个节点（避免断链）
        curr.Next = prev  //反转当前节点的指针，指向prev
        prev = curr       //prev移动到当前节点
        curr = temp       //curr移动到之前保存的下一个节点
    }
    return prev//go语言用头节点表示一个链表，链表没有单独的名字
}