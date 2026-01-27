package main

import (
	"fmt"
)	
func average(salary []int) float64 {
    a:=salary
	var i,j int
	for i=0;i<len(a);i++{
		for j=0;j<len(a)-1;j++{
			if a[j]>a[j+1]{
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	}
	var sum int
	for i=0;i<len(a);i++{
		sum+=(a[i])
	}
	sum=sum-(a[0])-(a[len(a)-1])
	return float64(sum)/float64(len(a)-2)
}

func main(){
	fmt.Println(average([]int{4000,3000,1000,2000}))
}
