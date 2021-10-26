package main

import (
	"fmt"
	"os"
)

func main() {
	os.Args = append(os.Args, "4")
	os.Args = append(os.Args, "2")
	if len(os.Args) < 3 {
		fmt.Print("incomplete ...")
		return
	}
	int1 := atoi(os.Args[1])
	int2 := atoi(os.Args[2])
	temp := int1
	answer := "X"
	if int1 > 26 || int1 == 0 || int2 == 0 {
		fmt.Print("not possible ...")
		return
	}
	N := int1 - 2
	if int1 > 2 {
		for i := 'A'; i < int32(N+65); i++ {
			answer += string(i)
			temp = (int(i) - 65) + 1
		}
	}
	answer += space(int1, int2)
	if int1 > 2 && int2 > 1 {
		for i := int32(temp + 65); i <= int32(N+temp+65)-1; i++ {
			answer += string(i)
		}
	}
	answer += "X"
	fmt.Print(answer)
}

func space(n1, n2 int) string {
	if n1 == 1 && n2 == 1 {
		answer := ""
		return answer
	}
	if n1 > 1 && n2 == 1 {
		answer := "X"
		return answer
	}
	if n1 == 1 && n2 > 1 {
		answer := "\n"
		for i := 1; i < n2-1; i++ {
			answer += "|\n"
		}
		return answer
	}
	if n1 > 1 && n2 > 1 {
		answer := "X\n"
		for i := 1; i < n2-1; i++ {
			answer += "|"
			for i := 1; i < n1-1; i++ {
				answer += " "
			}
			answer += "|\n"
		}
		answer += "X"
		return answer
	}
	return "\nX"
}

func atoi(s string) int {
	arrayStr := []rune(s)
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		if arrayStr[i] < '0' || arrayStr[i] > '9' {
			fmt.Print("incomplete ...")
			return 0
		} else {
			ans *= 10
			ans += int(arrayStr[i]) - '0'
		}
	}
	return ans
}
