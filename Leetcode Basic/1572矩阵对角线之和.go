package main

import "fmt"

func diagonalSum(mat [][]int) int {
	sum:=0
	if len(mat)%2==1{//奇
		for i:=0;i<len(mat);i++{//边长即行数，用len
			for j:=0;j<len(mat);j++{
				if i==j&&i+j!=len(mat)-1{//主对角线
					sum+=mat[i][j]
				}else if i+j==len(mat)-1{//副对角线
					sum+=mat[i][j]
				}
			}
		}
	}else if len(mat)%2==0{//偶数
		for i:=0;i<len(mat);i++{//边长即行数，用len
			for j:=0;j<len(mat);j++{
				if i==j{//主对角线
					sum+=mat[i][j]
				}else if i+j==len(mat)-1{//副对角线
					sum+=mat[i][j]
				}
			}
		}
	}
	return sum
}

func main(){
	fmt.Println(diagonalSum([][]int{{1,2,3},
            {4,5,6},
            {7,8,9}}))
}