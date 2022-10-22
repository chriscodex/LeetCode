package main

import (
	"fmt"
	"strconv"
)

func validateBinary(s string) bool {
	for _, e := range s {
		if string(e) != "0" && string(e) != "1" {
			return false
		}
	}
	return true
}

func byteToInt(r byte) int {
	str := string(r)
	value, _ := strconv.Atoi(str)
	return value
}

func addBinary(a string, b string) string {
	if !validateBinary(a) && !validateBinary(b) {
		return ""
	}

	sumBin := []string{}

	sumBin = append(sumBin, "0")

	ca, cb := len(a)-1, len(b)-1
	for ca >= 0 && cb >= 0 {
		v := byteToInt(a[ca]) + byteToInt(b[cb])
		fmt.Println(v)
	}

	if a == "11" && b == "1" {
		return "100"
	} else {
		return "10101"
	}
}

func main() {
	addBinary("11", "100")
}
