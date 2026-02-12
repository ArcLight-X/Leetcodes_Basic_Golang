package main

import("fmt")

func merge(intervals [][]int) [][]int {
    result:=make([][]int,0)
	
	//先按起点对数组排序
	for i:=0;i<len(intervals)-1;i++{
		for j:=0;j<len(intervals)-1;j++{
			if intervals[j][0]>intervals[j+1][0]{
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}
	//排序后，把每一组的终点和下一组起点比较
	current:=intervals[0]//当前区间
	for k:=1;k<len(intervals);k++{
		if current[1]>=intervals[k][0]{//本组终点>=下一组起点，合并(不用append,本质是把当前组的终点更新掉)
			current[1]=max(current[1],intervals[k][1])//终点取两组中较大的
		}else{//不合并
			result=append(result,current)//把当前区间加入结果
			current=intervals[k]//更新当前区间为下一个区间
		}
	}
	result=append(result,current)
	return result
}

func main(){
	fmt.Println(merge([][]int{{1,3},{2,6},{8,10},{15,18}}))
}