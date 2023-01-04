package kernel

import (
	"regexp"

	"shiza.io/src"
)

type Lexer struct {
	Code      string
	Pos       int64
	TokenList []Token
}

func NewLexer(codeContent string) *Lexer {
	return &Lexer{
		TokenList: []Token{},
		Code:      codeContent,
		Pos:       0,
	}
}

func (lexer *Lexer) RunLexicalAnalysis() {
	for lexer.getNextToken() {
	}
}

func (lexer *Lexer) getNextToken() bool {
	if lexer.Pos >= int64(len(lexer.Code)) {
		return false
	}
	for _, tokenType := range TokenQuerySlice {
		regexp := regexp.MustCompile("^" + tokenType.Reg)
		result := regexp.FindAllString(lexer.Code[int(lexer.Pos):], -1)
		if len(result) > 0 && len(result[0]) > 0 {
			if !tokenType.IsNotRequired {
				token := NewToken(tokenType, result[0], int64(lexer.Pos), int64(lexer.Pos))
				lexer.TokenList = append(lexer.TokenList, token)
			}
			lexer.Pos += int64(len(result[0]))
			return true
		}
	}
	src.LogLexerError(lexer.Pos, lexer.Code[lexer.Pos:lexer.Pos+15], true)
	return false
}
