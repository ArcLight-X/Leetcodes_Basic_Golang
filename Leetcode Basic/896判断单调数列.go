package main
import (
	"fmt"
)

func isMonotonic(nums []int) bool {
    a:=nums
	in:=true//假设是递增的
	de:=true//假设是递减的
	for i:=0;i<len(a)-1;i++{//有i+1，防止溢出循环到len-1
		if a[i]<a[i+1]{
			in=false//任两项大小不对，不递增
		}else if a[i]>a[i+1]{
			de=false//任两项大小不对，不递减
		}
	}
	return in||de   //或运算，有一个true就true                                                                                                                                                                                                                                                                                                                                                    
}

func main(){
	fmt.Println(isMonotonic([]int{6,5,4,4}))
}