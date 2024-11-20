package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// define the 3 repetition character constants here
const (
	prefix       = "^"
	suffix       = "$"
	questionMark = "?"
	asterisk     = "*"
	plusSign     = "+"
	backSlash    = "\\"
)

func main() {
	regex, text := spitRegAndText()

	fmt.Printf("%t\n", checkRegMatch(regex, text))
}

func spitRegAndText() (regex, text string) {
	var inputStr string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr = scanner.Text()
	inputSlc := strings.Split(inputStr, "|")
	regex = inputSlc[0]
	text = strings.Join(inputSlc[1:], "")

	if strings.Contains(regex, backSlash) {
		regex = strings.Replace(regex, backSlash, "", 1)
	} else {
		r, isOk := checkRepetition(regex, text)
		if !isOk {
			return
		}
		regex = r
	}

	return regex, text
}

func checkFirstChar(regex, text string) bool {
	return regex[0] == text[0] || regex[0] == '.'
}

func checkRepetition(regex, text string) (r string, isOk bool) {
	var repCharIndex, oCount int

	switch {
	// THIS IS SORTED OUT - NO TOUCHY
	case strings.Contains(regex, questionMark):
		repCharIndex = strings.Index(regex, questionMark)

		temp := strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], "")
		if regex[repCharIndex-1] == '.' {
			if text == temp {
				r, isOk = temp, true
			} else {
				r, isOk = strings.ReplaceAll(regex, questionMark, ""), true
			}
		} else {
			temp = strings.ReplaceAll(regex, questionMark, "")
			if strings.Count(text, string(regex[repCharIndex-1])) > 1 {
				isOk = false
			} else {
				if strings.Count(text, string(regex[repCharIndex-1])) == 0 {
					r = strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], "")
				} else {
					r = temp
				}
				isOk = true
			}
		}

	// THIS ONE AS WELL!!
	case strings.Contains(regex, asterisk):
		repCharIndex = strings.Index(regex, asterisk)
		temp := strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], "")

		if regex[repCharIndex-1] == '.' {
			if text != temp {
				r = strings.ReplaceAll(regex, asterisk, "")
			} else {
				r = temp
			}
			isOk = true
		} else {
			temp = strings.ReplaceAll(regex, asterisk, "")
			if strings.Count(text, string(regex[repCharIndex-1])) >= 1 {
				oCount = strings.Count(text, string(text[repCharIndex-1]))
				temp = strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], strings.Repeat(string(text[repCharIndex-1]), oCount))
				if !strings.Contains(temp, text) {
					isOk = false
				} else {
					r, isOk = temp, true
				}
			} else {
				r, isOk = strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], ""), true
			}
		}

	case strings.Contains(regex, plusSign):
		repCharIndex = strings.Index(regex, plusSign)

		temp := strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], "")
		if regex[repCharIndex-1] == '.' {
			if strings.Contains(temp, text) {
				isOk = false
			} else {
				isOk = true
			}
			r = strings.ReplaceAll(regex, plusSign, "")
		} else {
			temp = strings.ReplaceAll(regex, plusSign, "")
			if strings.Count(text, string(regex[repCharIndex-1])) >= 1 {
				oCount = strings.Count(text, string(text[repCharIndex-1]))
				temp = strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], strings.Repeat(string(text[repCharIndex-1]), oCount))
				r, isOk = temp, true
			} else {
				isOk = false
			}
		}

	default:
		r, isOk = regex, true
	}

	return r, isOk
}

func checkRegMatch(regex, text string) bool {
	if len(regex) == 0 || regex == text {
		return true
	}
	if len(text) == 0 {
		return false
	}

	if strings.HasPrefix(regex, prefix) {
		if strings.HasSuffix(regex, suffix) {
			if !checkFirstChar(regex[1:], text) {
				return strings.EqualFold(regex[1:len(regex)-1], text)
			}
			return checkRegMatch(regex[2:len(regex)-1], text)
		} else {
			if !checkFirstChar(regex[1:], text) {
				return strings.HasPrefix(text, regex[1:])
			}
			return checkRegMatch(regex[2:], text[1:])
		}
	}

	if strings.HasSuffix(regex, suffix) {
		if !checkFirstChar(regex, text) {
			return strings.HasSuffix(text, regex[:len(regex)-1])
		}
		return checkRegMatch(regex[1:len(regex)-1], text)
	} else {
		if !checkFirstChar(regex, text) {
			return strings.HasSuffix(text, regex)
		}
		return checkRegMatch(regex[1:], text[1:])
	}
}
