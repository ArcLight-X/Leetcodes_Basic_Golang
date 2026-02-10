package main

import("fmt")
//固定窗口长度为p的长度+字符计数
func findAnagrams(s string, p string) []int {
	need:=make([]int,26)//记录要求的数组的字符频次
	for i:=0;i<len(p);i++{
		need[p[i]-'a']++
	}
	ans:=[]int{}

	isEqual := func(arr1, arr2 []int) bool {//辅助函数，判断两个切片是否相等
        for i:=0;i<26;i++{
            if arr1[i]!=arr2[i] {
                return false
            }
        }
        return true
    }
    
	for i:=0;i<=len(s)-len(p);i++{//开始滑动窗口
		window:=make([]int, 26)
		for j:=i;j<i+len(p);j++{
			window[s[j]-'a']++//计算当前窗口的字符频次
		}
		// 检查当前窗口是否匹配
        if isEqual(window, need) {
            ans = append(ans, i)//添加匹配的起始位置
        }
		
	}
	return ans
}

func main(){
	fmt.Println(findAnagrams("cbaebabacd","abc"))
}