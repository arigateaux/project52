package funnel

import (
	"strings"
)

// two strings of words to find similarities by removing one letter

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return strings.Replace(string(out), " ", "", 1)
}

func Funnel(first, second string) bool {
	for i := 0; i < len(first); i++ {
		// chr := string(first[i])
		newFirst := replaceAtIndex(first, ' ', i)
		if newFirst == second {
			return true
		}
		// fmt.Println(chr, newFirst)
	}
	return false
}
