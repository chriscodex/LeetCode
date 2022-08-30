package main

import (
	"fmt"
)

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

var val int

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

func total(arr []int, cont int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 2 {
		return compareTwo(arr[0], arr[1])
	}
	if cont == len(arr)-1 {
		val += arr[cont]
		return val
	}
	if cont == len(arr)-2 {
		fmt.Println(cont)
		val += compareTwo(arr[cont], arr[cont+1])
		return val
	}
	if arr[cont] >= arr[cont+1] {
		val += arr[cont]
		cont++
		return total(arr, cont)
	}
	if arr[cont] < arr[cont+1] {
		val += arr[cont+1] - arr[cont]
		cont += 2
		return total(arr, cont)
	}
	return val
}
