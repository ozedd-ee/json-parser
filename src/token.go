package src

const (
	QUOTE      = `"`
	LEFTBRACE  = "{"
	RIGHTBRACE = "}"
	LEFTBRACKET = "["
	RIGHTBRACKET = "]"
	WHITESPACE = " "
	COMMA = ","
	COLON = ":"
)

var STRUCTURAL_TOKENS = []string{"{", "}", ":", ",", "[", "]"}

var VALID_NUMBER_CHAR = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", "."}

const TRUE_LEN = 4
const FALSE_LEN = 5
const NIL_LEN = 3

type token struct {
	value interface{}
	position int
}

type node struct {
	value interface{}
}
