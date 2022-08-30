package main

func getThreeFirst(s string) *[]string {
	arr := []string{}
	str := ""
	for i := 1; i < len(s)+1; i++ {
		if i%3 == 0 {
			str += string(s[i-1])
			arr = append(arr, str)
			str = ""
		} else {
			str += string(s[i-1])
		}
	}
	if str != "" {
		arr = append(arr, str)
	}
	return &arr
}

func compareTwo(a, b int) int {
	numeral := 0
	if a == b {
		numeral += 2 * a
	}
	if a > b {
		numeral += a + b
	}
	if a < b {
		numeral += b - a
	}
	return numeral
}

func strToNum(s string) []int {
	arr := []int{}
	for _, e := range s {
		switch string(e) {
		case "M":
			arr = append(arr, 1000)
		case "D":
			arr = append(arr, 500)
		case "C":
			arr = append(arr, 100)
		case "L":
			arr = append(arr, 50)
		case "X":
			arr = append(arr, 10)
		case "V":
			arr = append(arr, 5)
		case "I":
			arr = append(arr, 1)
		}
	}
	return arr
}
