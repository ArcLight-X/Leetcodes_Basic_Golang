package main

import("fmt")

func repeatedSubstringPattern(s string) bool {
	//i为可能子串长度
    for i:=1;i<=len(s)/2;i++{//子串最多是原来的一半长，而且是原长约数
		if len(s)%i==0{
			match:=true
			for j:=i; j<len(s); j+=i{
            //对比当前段和子串是否一致
            	if s[j:j+i]!=s[:i]{
                	match = false
                	break
				}
			}
			if match{
				return true
			}
		}
	}
	return false
}

func main(){
	fmt.Println(repeatedSubstringPattern("abab"))
}