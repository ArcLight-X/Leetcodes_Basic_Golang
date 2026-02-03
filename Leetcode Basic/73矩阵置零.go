package main

import("fmt")

func setZeroes(matrix [][]int)  {
	rows,cols:=len(matrix),len(matrix[0])
	//一定要先记录0的位置，不然边查找边修改会新增0的位置
	zerorows:=make([]bool,rows)
	zerocols:=make([]bool,cols)
	for i:=0;i<rows;i++{
		for j:=0;j<cols;j++{
			if matrix[i][j]==0{
				zerorows[i]=true
				zerocols[j]=true
			}
		}
	}
	//记录完成后按照记录开始修改
	for i:=0;i<rows;i++{
		if zerorows[i]{
			for j:=0;j<cols;j++{
				matrix[i][j]=0
			}
		}
	}
	for j:=0;j<cols;j++{
		if zerocols[j]{
			for i:=0;i<rows;i++{
				matrix[i][j]=0
			}
		}
	}
	return
	

}

func main(){
	fmt.Println(setZeroes([][]int{{1,2,3},{4,0,6},{7,8,9}}))
}