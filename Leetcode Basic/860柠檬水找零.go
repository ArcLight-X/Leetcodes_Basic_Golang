package main

import (
	"fmt"
)	

func lemonadeChange(bills []int) bool {
	charge5,charge10:=0,0
    for _,bill :=range bills{//只要value
		switch bill {
		case 5:
			charge5++
		case 10:
				if charge5==0{
					return false//没有5元就找不出了
				}
			charge5--
			charge10++//有的话就找5元，拿10元
		case 20:
			if charge10>0&&charge5>0{
				charge5--
				charge10--//优先把10元用了，10+5找零
			}else if charge10==0&&charge5>=3{
				charge5=charge5-3//没有就只能找3个5元，得有3个以上5元
			}else{
				return false
			}
		default:
			return false

	}
	}
	return true//执行到这才能true
}

func main(){
	fmt.Println(lemonadeChange([]int{5,5,5,10,20}))
}