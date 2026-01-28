package main

import "fmt"

func myPow(x float64, n int) float64 {
    // 处理 n 为负数的情况：x^-n = 1/(x^n)
    if n < 0 {
        x = 1 / x
        n = -n
    }
    result := 1.0
    current := x
    // 快速幂核心逻辑
    for n > 0 {
        // 如果当前指数是奇数，就把当前值乘到结果里
        if n%2 == 1 {
            result *= current
        }
        // 底数平方
        current *= current
        // 指数除以 2（向下取整）
        n = n / 2
    }
    return result
}

func main() {
    fmt.Println(myPow(2.0, 10))
    fmt.Println(myPow(2.0, -2))
    fmt.Println(myPow(3.0, 3)) 
}