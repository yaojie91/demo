package leet524

import "fmt"

func main() {
	var s = "abpcplea"
	var d = []string{"a", "b", "c"}
	fmt.Println("----------", findLongestWord(s, d))
}

func findLongestWord(s string, d []string) string {
	longest := ""
	for _, v := range d {
		if len(v) > len(longest) || (len(v) == len(longest) && v < longest) {
			if isSubStr(s, v) {
				longest = v
			}
		}
	}
	return longest
}

func isSubStr(s string, t string) bool {
	index1, index2 := 0, 0
	for index1 < len(s) && index2 < len(t) {
		if s[index1] == t[index2] {
			index2++
		}
		index1++
	}
	if index2 == len(t) {
		return true
	}
	return false
}
