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
	case strings.Contains(regex, questionMark):
		repCharIndex = strings.Index(regex, questionMark)
		temp := strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], "")
		if regex[repCharIndex-1] == '.' {
			if *text != temp {
				r = strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], string(tVal[repCharIndex-1]))
			} else {
				r = temp
			}
			//isPresent, isOk = true, true
			isPresent, isOk = true, true
		} else {
			if strings.Count(*text, string(regex[repCharIndex-1])) == 1 {
				r, isOk = strings.ReplaceAll(regex, questionMark, ""), true
			} else if !strings.Contains(*text, string(regex[repCharIndex-1])) {
				r, isOk = temp, true
			} else {
				r, isOk = strings.ReplaceAll(regex, questionMark, ""), false
			}
			isPresent = true
		}

	case strings.Contains(regex, asterisk):
		repCharIndex = strings.Index(regex, asterisk)
		temp := strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], "")

		if regex[repCharIndex-1] == '.' {
			if *text != temp {
				r = strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], string(tVal[repCharIndex-1]))
				*text = r
			} else {
				r = temp
			}
			//isPresent, isOk = true, true
			isOk = true
		} else {
			temp = strings.ReplaceAll(regex, asterisk, "")
			if len(tVal) >= len(temp) && regex[repCharIndex-1] != tVal[repCharIndex-1] {
				isOk = false
			} else if len(tVal) < len(temp) && regex[repCharIndex-1] != tVal[repCharIndex-1] {
				r, isOk = strings.ReplaceAll(regex, regex[repCharIndex-1:repCharIndex+1], ""), true
			} else {
				r, isOk = temp, true
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
				isOk = true
			}
			r = temp
		} else {
			temp = strings.ReplaceAll(regex, plusSign, "")
			if strings.Contains(tVal, string(regex[repCharIndex-1])) {
				isOk = true
			} else {
				isOk = false
			}
			r = temp
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

	// this could be part of the function for checking repetition chars
	if ((regex == ".?" || regex == ".+") && len(text) != 0) || regex == ".*" {
		return true
	}

	r, isPresent, isOk := checkRepetition(regex, &text)

	if isPresent && !isOk {
		return false
	}

	if !checkFirstChar(r, text) {

		if strings.HasPrefix(r, prefix) {
			if strings.HasSuffix(r, suffix) {
				// r, isPresent, _ = checkRepetition(regex[1:len(regex)-1], &text)
				if !isPresent {
					return strings.EqualFold(regex[1:len(regex)-1], text)
				} else {
					return /*checkRegMatch(r, text[1:]) */ strings.EqualFold(r[1:len(r)-1], text)
				}

			}
			if checkFirstChar(r[1:], string(text[0])) {
				r, isPresent, _ = checkRepetition(regex[1:], &text)
				if !isPresent {
					return strings.HasPrefix(text, regex[1:])
				}
				return checkRegMatch(r[1:], text[1:])
			}
		}

		if strings.HasSuffix(r, "$") {
			// r, _, _ = checkRepetition(regex[:len(regex)-1], text)
			return strings.HasSuffix(text, r[:len(r)-1])
		}

		return strings.HasSuffix(text, r)
	}

	//if the first is OK AND has suffix
	if strings.HasSuffix(regex, "$") {
		r, isPresent, _ = checkRepetition(regex[:len(regex)-1], &text)
		if !isPresent {
			return checkRegMatch(regex[:len(regex)-1], text[1:]) || strings.HasSuffix(text, regex[:len(regex)-1])
		}
		return checkRegMatch(r[1:], text[1:]) || strings.HasSuffix(text, r)
	} else {
		if !isPresent && len(r) > len(text) {
			return false
		}
		return checkRegMatch(r[1:], text[1:])
	}

}
