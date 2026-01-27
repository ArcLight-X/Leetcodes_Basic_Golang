package main
import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
    m, n := len(word1), len(word2)
    ans := make([]byte, 0, m+n)
	//准备一个slice切片，长度len目前为0，容量cap为m+n确保够大
    for i := 0; i < m || i < n; i++ {// 保证两个数组都循环，循环次数是较大的那一个
        if i < m {
            ans = append(ans, word1[i])//先执行，确保word1的先插入
        }
        if i < n {
            ans = append(ans, word2[i])
        }
    }
    return string(ans)
}

func main(){
	fmt.Println(mergeAlternately("ab","pqrs"))
}