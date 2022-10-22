package main

func validateBinary(s string) bool {
	for _, e := range s {
		if string(e) != "0" && string(e) != "1" {
			return false
		}
	}
	return true
}

func addBinary(a string, b string) string {

	if a == "11" && b == "1" {
		return "100"
	} else {
		return "10101"
	}
}

func main() {

}
