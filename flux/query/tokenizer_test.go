package query

import (
	"fmt"
	"strconv"
	"testing"
)

func print_tokens(tokens []string) {
	for _, a := range tokens {
		fmt.Println( a ) // debug
	}
}

func TestTokenizerWhenEmpty_ExpectNoError(t *testing.T) {
	tokens, _ := tokenize("")

	if len(tokens) != 0 {
		t.Error()
	}
}

func TestTokenizerSimple_ExpectNoError(t *testing.T) {
	tokens, _ := tokenize(" hello my name is ")

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
	tokens, _ := tokenize(" name is \"afshin the clown\" wow")

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
	tokens, _ := tokenize(" name is \"afshin \\\"the clown\\\" so \" wow")

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
	tokens, _ := tokenize("age = 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
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

func TestTokenizerWithSimpleEquals_noSpaces_ExpectError(t *testing.T) {
	_, err := tokenize("age=45")

	if err == nil {
		t.Error("expected error, but got no error")
	}
}

func TestTokenizerWithSimpleEquals2_noSpaces_ExpectError(t *testing.T) {
	_, err := tokenize("age= 45")

	if err == nil {
		t.Error("expected error, but got no error")
	}
}

func TestTokenizerWithSimpleEquals3_noSpaces_ExpectError(t *testing.T) {
	_, err := tokenize("age =45")

	if err == nil {
		t.Error("expected error, but got no error")
	}
}

func TestTokenizerWithoOPS_1(t *testing.T) {
	tokens, _ := tokenize("age = 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
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

func TestTokenizerWithoOPS_2(t *testing.T) {
	tokens, _ := tokenize("age != 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "!=" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}

func TestTokenizerWithoOPS_3(t *testing.T) {
	tokens, _ := tokenize("age <= 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "<=" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}
func TestTokenizerWithoOPS_4(t *testing.T) {
	tokens, _ := tokenize("age < 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "<" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}
func TestTokenizerWithoOPS_5(t *testing.T) {
	tokens, _ := tokenize("age > 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != ">" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}
func TestTokenizerWithoOPS_6(t *testing.T) {
	tokens, _ := tokenize("age >= 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != ">=" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}
func TestTokenizerWithoOPS_7(t *testing.T) {
	tokens, _ := tokenize("age StartsWith 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "StartsWith" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}
func TestTokenizerWithoOPS_8(t *testing.T) {
	tokens, _ := tokenize("age NotStartsWith 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "NotStartsWith" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}
func TestTokenizerWithoOPS_9(t *testing.T) {
	tokens, _ := tokenize("age EndsWith 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "EndsWith" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}
func TestTokenizerWithoOPS_10(t *testing.T) {
	tokens, _ := tokenize("age NotEndsWith 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "NotEndsWith" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}
func TestTokenizerWithoOPS_11(t *testing.T) {
	tokens, _ := tokenize("age Contains 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "Contains" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}
func TestTokenizerWithoOPS_12(t *testing.T) {
	tokens, _ := tokenize("age NotContains 45")

	if len(tokens) != 3 {
		t.Error("expected 3 but got " + strconv.Itoa(len(tokens)))
	}

	if tokens[0] != "age" {
		t.Error()
	}

	if tokens[1] != "NotContains" {
		t.Error()
	}

	if tokens[2] != "45" {
		t.Error()
	}
}
