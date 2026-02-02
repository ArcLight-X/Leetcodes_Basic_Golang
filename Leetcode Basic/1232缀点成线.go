package main

import("fmt")

func checkStraightLine(coordinates [][]int) bool {
	x:=make([]int,len(coordinates))
	y:=make([]int,len(coordinates))//切片提前初始化，防止越界
    for i,pos:=range coordinates{
		x[i]=pos[0]
		y[i]=pos[1]
	}
	for j:=2;j<len(coordinates);j++{//用乘法判断，避免除0，P0P1斜率=P0PX斜率
		if (y[1]-y[0])*(x[j]-x[0])!=(y[j]-y[0])*(x[1]-x[0]){
			return false
		}
	}
	return true
}

func main(){
	fmt.Println(checkStraightLine([][]int{{1,2},{2,3},{3,4},{4,5},{5,6},{6,7}}))
}