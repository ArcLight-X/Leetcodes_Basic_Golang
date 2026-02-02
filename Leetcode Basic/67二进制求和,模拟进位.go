package main

import "fmt"

func addBinary(a string, b string) string {
    res := ""
    i,j := len(a)-1, len(b)-1//从最低位开始
    carry:=0//进位

    //从低位到高位逐位相加
    for i >= 0||j >= 0||carry>0{//最高位不是0，说明还没到最高位，继续相加
        sum:=carry//sum是当前位结果，等于carry+a[当前位]+b[当前位]
        if i>=0 {
            sum += int(a[i] - '0')//字符串转整数
            i--
        }
        if j>=0 {
            sum += int(b[j] - '0')
            j--
        }
        // 当前位的结果是 sum % 2，二进制，1就是1，2变成0
        res=fmt.Sprintf("%d", sum%2) + res
        // 进位是 sum / 2
        carry=sum/2//二进制，sum等于2说明有1进位
    }
    return res
}

func main(){
	fmt.Println(addBinary("1010","1011"))
}