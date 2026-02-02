package main

import("fmt")

func addBinary(a string, b string) string {
    // 先把字符串转成整数
    numA,numB := 0,0
    // 字符串转整数
    for _, c:=range a {
		//101为例：第一步：0*2+1=1
		//第二步：1*2+0=2
		//第三步：2*2+1=5
        numA=numA*2+int(c-'0')
    }
    for _, c:=range b {
        numB=numB*2+int(c-'0')
    }

    //位运算实现加法
    for numB != 0 {
        // 1. 算本位和（异或）
        sum := numA ^ numB
        // 2. 算进位（与 + 左移1位）
        carry := (numA & numB) << 1
        // 3. 更新a为本位和，b为进位，直到进位为0
        numA, numB = sum, carry
    }

    // 整数转回二进制字符串
    if numA == 0 {
        return "0"
    }
    res := ""//除2取余法
    for numA > 0 {
        res = fmt.Sprintf("%d", numA%2) + res
        numA =numA/2
    }
    return res
}

func main(){
	fmt.Println(addBinary("1010","1011"))
}