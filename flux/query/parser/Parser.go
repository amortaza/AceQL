package parser

import (
	"github.com/amortaza/aceql/flux/logger"
	"github.com/amortaza/aceql/flux/node"
)


func Parse( encodedQuery string ) (node.Node, error) {
	stack := newGenericNodeStack()

	tokens := Tokenize(encodedQuery)

	for _, token := range tokens {
		stack.Push( token )
		stackReduce(stack)
	}

	return nil, nil
}

func stackReduce(stack *StringStack) {

}

func Tokenize(s string) []string {
	tokens := make([]string, 0)
	token := ""
	state := 0

	for _, cur := range s {
		if state == 0 {
			if cur == ' ' {
				continue
			} else if isUniToken(cur) {
				tokens = append(tokens, string(cur))
			} else if cur == '"' {
				token = "\""
				state = 2
				continue
			} else {
				// letter
				token += string(cur)
				state = 1
			}
		} else if state == 1 {
			if cur == ' ' {
				tokens = append(tokens, token)
				token = ""
				state = 0
			} else if isUniToken(cur) {
				tokens = append(tokens, token)
				tokens = append(tokens, string(cur))
				token = ""
				state = 0
			} else if cur == '"'{
				tokens = append(tokens, token)
				token = "\""
				state = 2
			} else {
				// letter
				token += string(cur)
			}
		} else if state == 2 {
			if cur == '"' {
				token += string(cur)
				tokens = append(tokens, token)
				token = ""
				state = 0
			} else if cur == '\\' {
				token += string(cur)
				state = 3
			} else {
				// anything else
				token += string(cur)
			}
		} else if state == 3 {
			token += string(cur)
			state = 2
		}
	}

	if token != "" {
		tokens = append(tokens, token)
	}

	if state != 0 && state != 1 && state != 3 {
		logger.Error("failed to parse encoded query, see ---" + s + "---", "Parser()")
		return make([]string, 0)
	}

	return tokens
}

func isUniToken(t rune) bool {
	result := t == '.' || t == '(' || t == ')'
	return result
}

