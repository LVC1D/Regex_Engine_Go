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
	wildCard     = "."
	backSlash    = "\\"
)

func main() {
	// write your code here
	regex, text := spitRegAndText()

	// we format-print the boolean result from the function call
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

	return regex, text
}

func checkFirstChar(regex, text string) bool {
	return regex[0] == text[0] || regex[0] == '.'
}

func checkRepetition(regex string, text *string) (r string, isPresent, isOk bool) {
	var repCharIndex int

	tVal := *text

	switch {
	// THIS IS SORTED OUT - NO TOUCHY
	case strings.Contains(regex, questionMark):
		repCharIndex = strings.Index(regex, questionMark)
		temp := strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], "")

		if regex[repCharIndex-1] == '.' {
			if tVal != temp {
				r, isOk = strings.ReplaceAll(regex, questionMark, ""), true
			} else {
				r, isOk = temp, true
			}
			isPresent = true
		} else {
			temp = strings.ReplaceAll(regex, questionMark, "")
			if strings.Count(*text, string(regex[repCharIndex-1])) == 1 {
				r, isOk = temp, true
			} else if !strings.Contains(*text, string(regex[repCharIndex-1])) && len(*text) < len(temp) {
				r, isOk = strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], ""), true
			} else {
				r, isOk = temp, false
			}
			isPresent = true
		}

	// THIS ONE AS WELL!!
	case strings.Contains(regex, asterisk):
		repCharIndex = strings.Index(regex, asterisk)
		temp := strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], "")

		if regex[repCharIndex-1] == '.' {
			if *text != temp {
				r = strings.ReplaceAll(regex, asterisk, "")
			} else {
				r = temp
			}
			isOk = true
		} else {
			temp = strings.ReplaceAll(regex, asterisk, "")
			if regex[repCharIndex-1] != tVal[repCharIndex-1] && len(tVal) >= len(temp) {
				isOk = false
			} else if strings.Count(*text, string(regex[repCharIndex-1])) >= 1 {
				oCount := strings.Count(*text, string(tVal[repCharIndex-1]))
				temp = strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], strings.Repeat(string(tVal[repCharIndex-1]), oCount))
				r, isOk = temp, true
			} else {
				r, isOk = strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], ""), true
			}
		}
		isPresent = true

	case strings.Contains(regex, plusSign):
		repCharIndex = strings.Index(regex, plusSign)
		temp := strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], "")
		if regex[repCharIndex-1] == '.' {
			if *text == temp {
				isOk = false
			} else {
				r, isOk = strings.ReplaceAll(regex, plusSign, ""), true
			}
		} else {
			temp = strings.ReplaceAll(regex, plusSign, "")
			if strings.Count(tVal, string(regex[repCharIndex-1])) >= 1 {
				oCount := strings.Count(*text, string(tVal[repCharIndex-1]))
				temp = strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], strings.Repeat(string(tVal[repCharIndex-1]), oCount))
				r, isOk = temp, true
			} else {
				r, isOk = temp, false
			}
		}
		isPresent = true

	default:
		r, isPresent, isOk = regex, false, true
	}
	return r, isPresent, isOk
}

// The recursive way

func checkRegMatch(regex, text string) bool {
	if len(regex) == 0 || regex == text {
		return true
	}
	if len(text) == 0 {
		return false
	}
	r, isPresent, isOk := checkRepetition(regex, &text)

	if isPresent && !isOk {
		return false
	}

	isFirstOk := checkFirstChar(r, text)

	if strings.HasPrefix(r, prefix) {
		isFirstOk = checkFirstChar(r[1:], text)
		if isFirstOk {
			if strings.HasSuffix(r, suffix) {
				if !isPresent {
					return strings.EqualFold(text, r[1:len(r)-1])
				}
				return checkRegMatch(r[2:len(r)-1], text[1:])
			} else {
				return checkRegMatch(r[2:], text[1:])
			}
		} else {
			return strings.HasSuffix(text, r)
		}
	}

	if !strings.HasSuffix(r, suffix) {
		if !isFirstOk {
			return strings.HasSuffix(text, r)
		}
		return checkRegMatch(r[1:], text[1:])
	}
	return checkRegMatch(r[1:len(r)-1], text)
}
