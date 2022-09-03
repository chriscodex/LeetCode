package main

import "fmt"

func compareTwoStrings(first string, second string) string {
	if first == "" || second == "" {
		return ""
	}
	arr := []string{}
	var compare int
	if len(first) < len(second) {
		compare = len(first)
	} else {
		compare = len(second)
	}
	for i := 0; i < compare; i++ {
		if first[i] == second[i] {
			arr = append(arr, string(first[i]))
		} else {
			break
		}
	}
	s := ""
	for _, e := range arr {
		s += e
	}
	return s
}

func main() {
	strs := []string{"ab", "a"}
	lcp := longCommPref(strs)
	fmt.Println(lcp)
}
