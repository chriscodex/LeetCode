package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	word := []byte{}

	cprefix := make(map[int]byte)

	for _, e := range strs[0] {
		word = append(word, byte(e))
	}

	for j := 1; j < len(strs)-1; j++ {
		cont := 0
		for cont <= len(strs[j])-1 {
			if word[cont] == strs[j][cont] {
				cprefix[cont] = word[cont]
				cont++
			} else {
				delete(cprefix, cont)
				cont++
			}
		}
	}
	s := ""
	for _, e := range cprefix {
		s += string(e)
	}
	return s
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	lcp := longestCommonPrefix(strs)
	fmt.Println(lcp)
}
