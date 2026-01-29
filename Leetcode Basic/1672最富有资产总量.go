package main

import("fmt"
)

func maximumWealth(accounts [][]int) int {
	max:=0
    for i:=0;i<len(accounts);i++{//len(a)为行数
		sum:=0//每一行前先清零
		for j:=0;j<len(accounts[i]);j++{//len(a[i])为子切片长度，即列数
			sum+=accounts[i][j]
		}
		if sum>max{
			max=sum
		}
	}
	return max
}

func main(){
	fmt.Println(maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}