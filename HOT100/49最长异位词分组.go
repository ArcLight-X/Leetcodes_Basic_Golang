package main

import("fmt"
"sort"
)

func groupAnagrams(strs []string) [][]string {
    //优化:字母异位词排序后必相同
	//哈希表：排序后字符串为key,原字符串数组为value
	m:=make(map[string][]string,len(strs))
	for _,s:=range strs{//先切割原字符串数组
		b:=[]byte(s)//转成字符串切片
		sort.Slice(b,func(i,j int)bool{
			return b[i]<b[j]//ascii码升序排序,即abcd顺序
		})
		sorted:=string(b)//转回字符串
		m[sorted]=append(m[sorted],s)//按上面键值对逻辑加入哈希表
		//排序后相同的字符串即为同一组
	}
		result:=make([][]string,0,len(m))//创建结果数组
		for _,v:=range m{
			result=append(result,v)
		}
	return result
}

func main(){
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}