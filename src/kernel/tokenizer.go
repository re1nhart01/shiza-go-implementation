package kernel

import "fmt"

type TokenQuery struct {
	Name          string
	Reg           string
	IsNotRequired bool
}

type Token struct {
	Type        TokenQuery
	Content     string
	PositionCol int64
	PositionRow int64
}

func (token *Token) Print() {
	fmt.Println(token.Type.Name, token.Content, token.PositionCol, token.PositionRow)
}

func NewToken(tokenType TokenQuery, content string, col int64, row int64) Token {
	return Token{
		Type:        tokenType,
		Content:     content,
		PositionCol: col,
		PositionRow: row,
	}
}

var TokenQuerySlice = map[string]TokenQuery{
	"SEMICOLON": {
		Name: "SEMICOLON",
		Reg:  ";",
	},
	"FUNCTION": {
		Name: "FUNCTION",
		Reg:  `do [A-z0-9]+\s?\[]\s?`, //do ([A-z0-9]+)\s?\[]\s?\((?:[^}{]+|\{(?:[^}{]+|\{[^}{]*\})*\})*\)
	},
	"VARIABLE": {
		Name: "VARIABLE",
		Reg:  "define ",
	},
	"VARIABLE_IN_USE": {
		Name: "VARIABLE_IN_USE",
		Reg:  `@[a-zA-Z]+`,
	},
	"ASSIGN": {
		Name: "ASSIGN",
		Reg:  `\s?=\s?`,
	},
	"CONSTANT": {
		Name: "CONSTANT",
		Reg:  "constant ",
	},

	//Types
	"NUMBER_INTEGER8": {
		Name: "NUMBER8",
		Reg:  "[0-9]",
	},
	"NUMBER_INTEGER16": {
		Name:          "NUMBER16",
		Reg:           "some sort of integer",
		IsNotRequired: true,
	},
	"STRING": {
		Name: "STRING",
		Reg:  `"[^"]+"`,
	},
	"BOOL": {
		Name: "BOOL",
		Reg:  "(true|false)+",
	},
	"LPAR": {
		Name: "LEFT_PAR",
		Reg:  `\s?\(\s?`,
	},
	"RPAR": {
		Name: "RIGHT_PAR",
		Reg:  `\s?\)\s?`,
	},
	"PLUS": {
		Name: "PLUS",
		Reg:  `\s?\+\s?`,
	},
	"MINUS": {
		Name: "MINUS",
		Reg:  `\s?\-\s?`,
	},
	"MULTIPLY": {
		Name: "MULTIPLY",
		Reg:  `\s\*\s`,
	},
	"DIVIDE": {
		Name: "DIVIDE",
		Reg:  `\s?\/\s?`,
	},
	"DISPLAY": {
		Name: "DISPLAY",
		Reg:  `display`,
	},
	"FOREVER_LOOP": {
		Name: "FOREVER LOOP",
		Reg:  `forever\s?`,
	},
	"TIMEOUT": {
		Name: "TIMEOUT",
		Reg:  `sleep`,
	},
	"SPACE": {
		Name:          "SPACE",
		Reg:           `[ \n\t\r]`,
		IsNotRequired: true,
	},
}
