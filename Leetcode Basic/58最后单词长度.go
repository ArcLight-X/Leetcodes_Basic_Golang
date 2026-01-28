package main

import (
	"fmt"
)	
func lengthOfLastWord(s string) int {
    end:= len(s)-1
    for end>=0 && s[end] == ' '{//s[end]从后往前找空格
        end--
    }//遇到第一个非空格，循环就停止
    if end<0 {
        return 0//全是空格就返回0
    }
    s=s[:end+1]//截断末尾空格后的字符串

    //从截断后的字符串末尾往前找第一个空格，中间的字符数就是最后一个单词长度
    start := end
    for start>=0&&s[start]!=' '{
        start--//遇到第一个空格，循环就停止
    }

    return end-start
}

func main(){
	fmt.Println(lengthOfLastWord("   fly"))
}