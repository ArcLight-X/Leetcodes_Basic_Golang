package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle)==0{
		return 0
	}else if len(haystack)<len(needle){
		return -1
	}
    for i:=0;i<=len(haystack)-len(needle);i++{
		//要检查的长度肯定大于等于len(needle),同时也防止下面[i+j]越界
		match:=true
		for j:=0;j<len(needle);j++{
			if haystack[i+j]!=needle[j]{
		//这里i是起始位置，看haystack里i+j个是不是和needle里从j开始计数完全匹配
				match =false
				break
			}
		}
		if match{
			return i
		}
		}
		return -1
	}


func main(){
	fmt.Println(strStr("sadbutsad","sad"))
}