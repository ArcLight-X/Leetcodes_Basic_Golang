package main

import (
	"fmt"
)	

func tictactoe(moves [][]int) string {
    var A,B [][]int//开新的二维数组，A是a走的坐标，B是b....
	for id,row:=range moves{
		if(id+1)%2==1{//原二维数组的奇数个
				newrow:=make([]int, len(row))//新的slice，长度还是row，即2
				copy(newrow,row)//复制到newrow切片，防止干扰
				A=append(A,newrow)//加到A的行里
		}else if (id+1)%2==0{//偶数个，同理
				newrow:=make([]int, len(row))
				copy(newrow,row)
				B=append(B,newrow)
		}
	}
	board:=[3][3]int{}//3*3二维数组，模拟棋盘
	for _, pos := range A {
		row, col := pos[0], pos[1]
		if row >= 0 && row < 3 && col >= 0 && col < 3 {//防越界
			board[row][col] = 1//A走的算1
		}
	}
	
	for _, pos := range B {
		row, col := pos[0], pos[1]
		if row >= 0 && row < 3 && col >= 0 && col < 3 {
			board[row][col] = -1//B走的算-1
		}
	}
	for i:=0;i<3;i++{
		rowSum:=board[i][0]+board[i][1]+board[i][2]//横向
		if rowSum==3{//和为3则A胜
			return "A"
		}else if rowSum==-3{
			return "B"
		}
	}
	for j := 0; j < 3; j++ {
		colSum := board[0][j] + board[1][j] + board[2][j]//纵向
		if colSum == 3 {
			return "A"
		} else if colSum == -3 {
			return "B"
		}
	}
	diag1Sum := board[0][0] + board[1][1] + board[2][2]//主对角线
	if diag1Sum == 3 {
		return "A"
	} else if diag1Sum == -3 {
		return "B"
	}
	diag2Sum := board[0][2] + board[1][1] + board[2][0]//副对角线
	if diag2Sum == 3 {
		return "A"
	} else if diag2Sum == -3 {
		return "B"
	}
	for k:=0;k<3;k++{
		for l:=0;l<3;l++{
			if board[k][l]==0{//遍历棋盘，还有0，即还有空没下，游戏未结束
			return "Pending" 
		}
		}
	}
	return "Draw"//没空了且没人赢，平局
}

func main(){
	fmt.Println(tictactoe([][]int{{0,0},{2,0},{1,1},{2,1},{2,2}}))
}