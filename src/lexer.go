/*
Copyright © 2024 Emmanuel Ozeh  github.com/ozedd-ee

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package src

import (
	"log"
	"strconv"
)

func Lex(jsonString string) []token {
	var tokens []token
	var currentPos int

	//Variable copy of jsonString: To be updated across iterations
	s := jsonString

	for i := 0; i <= len(jsonString); i++ {
		if string(s[0]) == WHITESPACE {
			currentPos++
			s = s[1:]
		}
		if len(s) == 0 {
			break
		}
		for _, v := range STRUCTURAL_TOKENS {
			if len(s) == 0 {
				break
			}
			if string(s[0]) == v {
				tokens = append(tokens, token{Value: v, position: currentPos})
				currentPos++
				s = s[1:]
			}
		}
		var jsonToken token

		jsonToken, s = lexString(s, currentPos)
		if jsonToken.Value != nil {
			currentPos++
			tokens = append(tokens, jsonToken)
		}

		jsonToken, s = lexNumber(s, currentPos)
		if jsonToken.Value != nil {
			currentPos++
			tokens = append(tokens, jsonToken)
		}

		jsonToken, s = lexBool(s, currentPos)
		if jsonToken.Value != nil {
			currentPos++
			tokens = append(tokens, jsonToken)
		}

		jsonToken, s = lexNil(s, currentPos)
		if jsonToken.Value == nil {
			currentPos++
			tokens = append(tokens, jsonToken)
		}

		if len(s) == 0 {
			break
		}
	}
	return tokens
}

func lexString(s string, currentPos int) (token, string) {

	if len(s) == 0 {
		return token{Value: nil}, s
	}
	jsonString := ""
	if string(s[0]) == QUOTE {
		s = s[1:]
	} else {
		return token{Value: nil}, s
	}
	for _, v := range s {
		if string(v) == QUOTE {
			return token{Value: jsonString, position: currentPos}, s[len(jsonString)+1:]
		} else {
			jsonString += string(v)
		}
	}
	log.Fatal("Expected end-of-string quote")
	return token{Value: nil}, s
}

func lexNumber(s string, currentPos int) (token, string) {
	if len(s) == 0 {
		return token{Value: nil}, s
	}
	jsonNumberString := ""
	numChar := make(map[string]bool)

	for _, val := range VALID_NUMBER_CHAR {
		numChar[val] = true
	}
	for _, c := range s {
		if numChar[string(c)] {
			jsonNumberString += string(c)
		} else {
			break
		}
	}

	if len(jsonNumberString) == 0 {
		return token{Value: nil}, s
	}
	s = s[len(jsonNumberString):]
	for _, v := range jsonNumberString {
		if string(v) == "." {
			jsonFloat, err := strconv.ParseFloat(jsonNumberString, 32)
			if err != nil {
				log.Fatal(err)
			}
			return token{Value: jsonFloat, position: currentPos}, s
		}
	}
	jsonInt, err := strconv.Atoi(jsonNumberString)
	if err != nil {
		log.Fatal(err)
	}
	return token{Value: jsonInt, position: currentPos}, s
}

func lexBool(s string, currentPos int) (token, string) {
	arrayLen := len(s)
	if arrayLen == 0 {
		return token{Value: nil}, s
	}
	if arrayLen >= TRUE_LEN && string(s[:TRUE_LEN]) == "true" {
		return token{Value: true, position: currentPos}, s[TRUE_LEN:]
	} else if arrayLen >= FALSE_LEN && string(s[:FALSE_LEN]) == "false" {
		return token{Value: false, position: currentPos}, s[FALSE_LEN:]
	}
	return token{Value: nil}, s
}

func lexNil(s string, currentPos int) (token, string) {
	if len(s) == 0 {
		return token{Value: ""}, s
	}
	if len(s) >= NIL_LEN && string(s[:NIL_LEN]) == "nil" {
		return token{Value: nil, position: currentPos}, s[NIL_LEN:]
	}
	return token{Value: ""}, s
}
