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
	"fmt"
	"log"
)

func parseArray(tokens []token) (Node, []token) {
	t := tokens[0]
	jsonArray := []any{}

	if t.Value == RIGHTBRACKET {
		return Node{isLeaf: true, LeafValue: jsonArray}, tokens[1:]
	}
	ts := tokens
	for {
		var node Node
		node, ts = Parse(ts)
		if node.isLeaf {
			jsonArray = append(jsonArray, node.LeafValue)
		} else {
			jsonArray = append(jsonArray, node.Value)
		}
		t := ts[0]
		if t.Value == RIGHTBRACKET {
			return Node{isLeaf: true, LeafValue: jsonArray}, ts[1:]
		} else if len(ts) == 0 {
			log.Fatalf("Expected end-of-array bracket ']' at position: %v, got : %v", t.position, t.Value)
			break
		} else if t.Value != COMMA {
			log.Fatalf("Expected comma ',' at the end of object at position: %v, got: %v", t.position, t.Value)
		} else {
			ts = ts[1:]
		}
	}
	return Node{}, tokens
}

func parseObject(tokens []token) (Node, []token) {
	node := make(map[string]any)
	t := tokens[0]
	if t.Value == RIGHTBRACE {
		return Node{Value: make(map[string]any)}, tokens[1:]
	}
	ts := tokens // variable copy of tokens

	for {
		jsonKey := ts[0]
		_, ok := jsonKey.Value.(string)
		if ok {
			ts = ts[1:]
		} else {
			log.Fatalf("Expected string key, at position: %v, got: %s", jsonKey.position, jsonKey.Value)
		}
		if ts[0].Value != COLON {
			log.Fatalf("Expected colon between key and Value at position: %v, got: %s", ts[0].position, ts[0].Value)
		}

		var jsonValue Node
		jsonValue, ts = Parse(ts[1:])
		if jsonValue.isLeaf {
			node[fmt.Sprint(jsonKey.Value)] = jsonValue.LeafValue
		} else {
			node[fmt.Sprint(jsonKey.Value)] = jsonValue.Value
		}

		t := ts[0]
		if t.Value == RIGHTBRACE {
			return Node{Value: node}, ts[1:]
		}
		if t.Value != COMMA {
			log.Fatalf("Expected end-of-Value comma at position: %v, got %s", t.position, t.Value)
		}
		if len(ts) != 0 {
			ts = ts[1:]
		} else {
			log.Fatalf("Expected end-of-object bracket '}'")
			break
		}
	}
	return Node{}, []token{}
}

func Parse(tokens []token) (Node, []token) {
	t := tokens[0]
	if t.Value == LEFTBRACE {
		return parseObject(tokens[1:])
	} else if t.Value == LEFTBRACKET {
		return parseArray(tokens[1:])
	}
	return Node{isLeaf: true, LeafValue: t}, tokens[1:]
}
