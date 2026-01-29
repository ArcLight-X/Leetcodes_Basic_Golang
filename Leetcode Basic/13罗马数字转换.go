package main

import("fmt"
)

func romanToInt(s string) int {
	romanMap := map[byte]int{//建立键值对
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	n := len(s)
	for i := 0; i < n-1; i++ { // 遍历到倒数第二个字符,防止越界
		current := romanMap[s[i]]
		next := romanMap[s[i+1]]
		if current < next {
			total -= current // 减法规则（如 IV → -1 +5）
		}else{
			total += current // 加法规则（如 VI → 5+1）
		}
	}
	total += romanMap[s[n-1]]//把最后一个加上
	return total
}

func main(){
	fmt.Println(romanToInt("III"))
}