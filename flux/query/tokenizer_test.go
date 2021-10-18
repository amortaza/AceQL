package query

import (
	"fmt"
	"testing"
)

func debug(tokens []string) {
	for _, a := range tokens {
		fmt.Println( a ) // debug
	}
}

func TestTokenizerWhenEmpty_ExpectNoError(t *testing.T) {
	tokens := Tokenize("")

	if len(tokens) != 0 {
		t.Error()
	}
}

func TestTokenizerSimple_ExpectNoError(t *testing.T) {
	tokens := Tokenize(" hello my name is ")

	if len(tokens) != 4 {
		t.Error()
	}

	if tokens[0] != "hello" {
		t.Error()
	}

	if tokens[1] != "my" {
		t.Error()
	}

	if tokens[2] != "name" {
		t.Error()
	}

	if tokens[3] != "is" {
		t.Error()
	}
}

func TestTokenizerWithQuotedString_ExpectNoError(t *testing.T) {
	tokens := Tokenize(" name is \"afshin the clown\" wow")

	if len(tokens) != 4 {
		t.Error()
	}

	if tokens[0] != "name" {
		t.Error()
	}

	if tokens[1] != "is" {
		t.Error()
	}

	if tokens[2] != "\"afshin the clown\"" {
		t.Error()
	}

	if tokens[3] != "wow" {
		t.Error()
	}
}

func TestTokenizerWithDoubleQuotedString_ExpectNoError(t *testing.T) {
	tokens := Tokenize(" name is \"afshin \\\"the clown\\\" so \" wow")

	if len(tokens) != 4 {
		t.Error()
	}

	if tokens[0] != "name" {
		t.Error()
	}

	if tokens[1] != "is" {
		t.Error()
	}

	if tokens[2] != "\"afshin \\\"the clown\\\" so \"" {
		t.Error()
	}

	if tokens[3] != "wow" {
		t.Error()
	}
}

func TestTokenizerWithSimpleEquals_ExpectNoError(t *testing.T) {
	tokens := Tokenize("age = 45")

	if len(tokens) != 3 {
		t.Error()
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "=" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}

func TestTokenizerWithSimpleEquals_noSpaces_ExpectNoError(t *testing.T) {
	tokens := Tokenize("age=45")

	if len(tokens) != 3 {
		t.Error()
		return
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "=" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}

func TestTokenizerWithSimpleEquals2_noSpaces_ExpectNoError(t *testing.T) {
	tokens := Tokenize("age= 45")

	if len(tokens) != 3 {
		t.Error()
		return
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "=" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}
func TestTokenizerWithSimpleEquals3_noSpaces_ExpectNoError(t *testing.T) {
	tokens := Tokenize("age =45")

	if len(tokens) != 3 {
		t.Error()
		return
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "=" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}
