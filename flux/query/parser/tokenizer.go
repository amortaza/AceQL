package parser

import (
	"errors"
	"github.com/amortaza/aceql/flux/query"
)

func tokenize(s string) ([]string, error) {
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
		return nil, errors.New("(6) failed to parse encoded query, see ---" + s + "--- tokenize.Parser()")
	}

	// now we go through all the tokens
	// if a token contains an operation like '=', it better be the whole operation
	// that is, it cannot be something like 'age=50', since we MUST surround operations by space
	for _, token := range tokens {
		if query.IsEncodedOps(token) {
			continue
		}

		if query.ContainsEncodedOps(token) {
			return nil, errors.New("operators must be separated by space, see ---" + s + "---")
		}
	}

	return tokens, nil
}

func isUniToken(t rune) bool {
	result := t == '.' || t == '(' || t == ')'
	return result
}


