package main

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
		astNode, tokens = parse(tokens)
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
		jsonValue, ts = parse(ts[1:])
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

func parse(tokens []token) (node, []token) {
	t := tokens[0]
	if t.value == LEFTBRACE {
		return parseObject(tokens[1:])
	} else if t.value == LEFTBRACKET {
		return parseArray(tokens[1:])
	}
	return node{value: t}, tokens[1:]
}
