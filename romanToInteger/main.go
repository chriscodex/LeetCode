package main

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
