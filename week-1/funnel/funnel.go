package funnel

import (
	"bufio"
	"os"
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
		// chr := string(first[i])  // debug
		newFirst := replaceAtIndex(first, ' ', i)
		if newFirst == second {
			return true
		}
		// fmt.Println(chr, newFirst)  // debug
	}
	return false
}

func wordSearch(word string) string {
	file, _ := os.Open(`./funnel/enable1.txt`)
	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		w := s.Text()
		if w == word {
			return w
		}
	}

	return ""
}

func isSliceMissing(slice []string, str string) bool {
	for _, ele := range slice {
		if ele == str {
			return false // item exists; return false
		}
	}
	return true // item is missing; return true
}

func Bonus(str string) []string {
	words := []string{} // empty slice for word storage

	for i := 0; i < len(str); i++ {
		tmp := replaceAtIndex(str, ' ', i) // return new word with removed letter
		word := wordSearch(tmp)            // return word from dictionary
		if word == tmp {
			isMissing := isSliceMissing(words, word) // check to see if the word is missing in the list
			if isMissing {
				words = append(words, word) // add word to list if missing
			}
		}
	}
	return words
}
