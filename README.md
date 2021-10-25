# project_challenge_go_s2

Project Challenge GO B1

os.Args = append(os.Args, "1")
os.Args = append(os.Args, "2")
if len(os.Args) < 3 {
fmt.Println("incomplete ...")
return
}
str1 := os.Args[1]
str2 := os.Args[2]
X := 0
temp := 0

	arrayStr := []rune(str1)
	n := len(str1)
	int1 := 0
	for i := 0; i < n; i++ {
		if arrayStr[i] < '0' || arrayStr[i] > '9' {
			fmt.Println("not possible...")
			return
		} else {
			int1 *= 10
			int1 += int(arrayStr[i]) - '0'
		}
	}

	arrayStr = []rune(str2)
	n = len(str2)
	int2 := 0
	for i := 0; i < n; i++ {
		if arrayStr[i] < '0' || arrayStr[i] > '9' {
			fmt.Println("not possible...")
			return
		} else {
			int2 *= 10
			int2 += int(arrayStr[i]) - '0'
		}
	}
	if int1 == 0 || int2 == 0 {
		fmt.Println("not possible...")
		return
	}
	for i := 0; i < int2; i++ {
		if i == 0 || i == int2-1 {
			for i := 0; i < int1; i++ {
				if int1 > 26 {
					X = int1 / 26
					int1 = int1 % 26
				}
				if i != 0 && i != int1-1 || X > 0 {
					if X > 0 {
						for X > 0 {
							for i := 'A'; i <= 'Z'; i++ {
								fmt.Print(string(i))
								X--
								if X == 0 {
									temp = int(i)
								}
							}
						}
					}
					if int1 > 0 {
						for i := 65+temp; i <= int1+65; i++ {
							fmt.Print(string(rune(i)))
							X--
							int1--
							if int1 == 0 {
								break
							}
						}
					}
				} else {
					fmt.Print("X")
					if i == int1-1 {
						fmt.Print("\n")
					}
				}
			}
		} else {
			fmt.Print("\n")
			fmt.Print("|")
			for i := 0; i < int1-2; i++ {
				fmt.Print(" ")
			}
			fmt.Print("|")
			fmt.Print("\n")
		}
	}
}
