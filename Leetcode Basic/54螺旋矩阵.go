package main

import("fmt")

func spiralOrder(matrix [][]int) []int {
    ans:=[]int{}
	top,bottom:=0,len(matrix)-1//行标
	left,right:=0,len(matrix[0])-1//列标
	for top<=bottom&&left<=right{
		for i:=left;i<=right;i++{
			ans=append(ans,matrix[top][i])//最上一行从左到右
		}
		top++//最上一行已经添加
		for i:=top;i<=bottom;i++{
			ans=append(ans,matrix[i][right])//最右一列从上到下
		}
		right--
		if top <= bottom {//检验避免单行重复
		for i:=right;i>=left;i--{//最下面一行从右到左
			ans=append(ans,matrix[bottom][i])
		}
		bottom--
		}
		if left<=right{//检验避免单列重复
			for i:=bottom;i>=top;i--{
			ans=append(ans,matrix[i][left])//最左边一行从下到上
			}
		left++
		}
	}
	return ans
}

func main(){
	fmt.Println(spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9}}))
}