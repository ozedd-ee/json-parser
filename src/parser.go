/*
Copyright © 2024 Emmanuel Ozeh  github.com/Eke-Manuel

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
	"fmt"
	"log"
)

func parseArray(tokens []token) (node, []token) {
	t := tokens[0]
	jsonArray := []any{}

	if t.value == RIGHTBRACKET {
		return node{value: jsonArray}, tokens[1:]
	}
	for {
		var astNode node
		astNode, tokens = Parse(tokens)
		jsonArray = append(jsonArray, astNode.value)
		t := tokens[0]
		if t.value == RIGHTBRACKET {
			return node{value: jsonArray}, tokens[1:]
		} else if len(tokens) == 0 {
			log.Fatalf("Expected end-of-array bracket ']' at position: %v, got : %v", t.position, t.value)
			break
		} else if t.value != COMMA {
			log.Fatalf("Expected comma ',' at the end of object at position: %v, got: %v", t.position, t.value)
		} else {
			tokens = tokens[1:]
		}
	}
	return node{}, tokens
}

func parseObject(tokens []token) (node, []token) {
	jsonObject := make(map[string]any)
	t := tokens[0]
	if t.value == RIGHTBRACE {
		return node{value: jsonObject}, tokens[1:]
	}
	ts := tokens // variable copy of tokens

	for {
		jsonKey := ts[0]
		_, ok := jsonKey.value.(string)
		if ok {
			ts = ts[1:]
		} else {
			log.Fatalf("Expected string key, at position: %v, got: %s", jsonKey.position, jsonKey.value)
		}
		if ts[0].value != COLON {
			log.Fatalf("Expected colon between key and value at position: %v, got: %s", ts[0].position, ts[0].value)
		}
		var jsonValue node
		jsonValue, ts = Parse(ts[1:])
		jsonObject[fmt.Sprint(jsonKey.value)] = jsonValue.value

		t := ts[0]
		if t.value == RIGHTBRACE {
			return node{value: jsonObject}, ts[1:]
		}
		if t.value != COMMA {
			log.Fatalf("Expected end-of-object comma at position: %v, got %s", t.position, t.value)
		}
		if len(ts) != 0 {
			ts = ts[1:]
		} else {
			log.Fatalf("Expected end-of-object bracket '}'")
			break
		}
	}
	return node{value: ""}, []token{}
}

func Parse(tokens []token) (node, []token) {
	t := tokens[0]
	if t.value == LEFTBRACE {
		return parseObject(tokens[1:])
	} else if t.value == LEFTBRACKET {
		return parseArray(tokens[1:])
	}
	return node{value: t}, tokens[1:]
}
