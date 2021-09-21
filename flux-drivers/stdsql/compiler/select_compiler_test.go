package compiler

import (
	"fmt"
	"testing"

	"github.com/amortaza/aceql/flux/node"
)

func TestWhenNoWhereClause_ExpectNoError(t *testing.T) {
	s := NewSelectCompiler("u_user", nil)
	q, _ := s.Compile()

	if q != "SELECT * FROM u_user" {
		t.Error()
	}
}

func TestWhenEqual_ExpectNoError(t *testing.T) {
	equal := node.NewEquals(NewNodeCompiler())
	equal.Left = node.NewColumn("name", NewNodeCompiler())
	equal.Right = node.NewString("ace", NewNodeCompiler())

	s := NewSelectCompiler("u_user", equal)
	q, _ := s.Compile()

	if q != "SELECT * FROM u_user WHERE name = 'ace'" {
		fmt.Print(q)
		t.Error()
	}
}

func TestNodeWhenAnd_ExpectNoError(t *testing.T) {
	name := node.NewEquals(NewNodeCompiler())
	name.Left = node.NewColumn("name", NewNodeCompiler())
	name.Right = node.NewString("ace", NewNodeCompiler())

	age := node.NewEquals(NewNodeCompiler())
	age.Left = node.NewColumn("age", NewNodeCompiler())
	age.Right = node.NewNumber(44, NewNodeCompiler())

	and := node.NewAnd(NewNodeCompiler())
	and.Left = name
	and.Right = age

	s := NewSelectCompiler("u_user", and)
	q, _ := s.Compile()

	if q != "SELECT * FROM u_user WHERE ( name = 'ace' AND age = 44.000000 )" {
		fmt.Print(q)
		t.Error()
	}
}

func TestNodeWhenAndOr_ExpectNoError(t *testing.T) {
	name := node.NewEquals(NewNodeCompiler())
	name.Left = node.NewColumn("name", NewNodeCompiler())
	name.Right = node.NewString("ace", NewNodeCompiler())

	age := node.NewEquals(NewNodeCompiler())
	age.Left = node.NewColumn("age", NewNodeCompiler())
	age.Right = node.NewNumber(44, NewNodeCompiler())

	and := node.NewAnd(NewNodeCompiler())
	and.Left = name
	and.Right = age

	or := node.NewOr(NewNodeCompiler())
	or.Left = and
	or.Right = age

	s := NewSelectCompiler("u_user", or)
	q, _ := s.Compile()

	if q != "SELECT * FROM u_user WHERE ( ( name = 'ace' AND age = 44.000000 ) OR age = 44.000000 )" {
		fmt.Print(q)
		t.Error()
	}
}

func TestWhenInStringClause_ExpectNoError(t *testing.T) {
	in := node.NewIn(NewNodeCompiler())
	in.Left = node.NewColumn("name", NewNodeCompiler())
	in.Right = node.NewStringList([]string{"ace", "clown", "reek"}, NewNodeCompiler())

	s := NewSelectCompiler("u_user", in)
	q, _ := s.Compile()

	if q != "SELECT * FROM u_user WHERE name IN [ 'ace', 'clown', 'reek' ]" {
		fmt.Print(q)
		t.Error()
	}
}

func TestWhenInNumberClause_ExpectNoError(t *testing.T) {
	in := node.NewIn(NewNodeCompiler())
	in.Left = node.NewColumn("name", NewNodeCompiler())
	in.Right = node.NewNumberList([]int{1, 2, 3, 4, 5}, NewNodeCompiler())

	s := NewSelectCompiler("u_user", in)
	q, _ := s.Compile()

	if q != "SELECT * FROM u_user WHERE name IN [ 1, 2, 3, 4, 5 ]" {
		fmt.Print(q)
		t.Error()
	}
}

func TestNodeWhenStartsWith_ExpectNoError(t *testing.T) {
	colnode := node.NewColumn("name", NewNodeCompiler())
	startsWith := node.NewStartsWith(NewNodeCompiler())

	startsWith.Left = colnode
	startsWith.Right = node.NewString("ace", NewNodeCompiler())

	s := NewSelectCompiler("u_user", startsWith)
	q, _ := s.Compile()

	if q != "SELECT * FROM u_user WHERE name LIKE 'ace%'" {
		fmt.Print(q)
		t.Error()
	}
}

func TestNodeWhenContains_ExpectNoError(t *testing.T) {
	colnode := node.NewColumn("name", NewNodeCompiler())
	contains := node.NewContains(NewNodeCompiler())

	contains.Left = colnode
	contains.Right = node.NewString("ace", NewNodeCompiler())

	s := NewSelectCompiler("u_user", contains)
	q, _ := s.Compile()

	if q != "SELECT * FROM u_user WHERE name LIKE '%ace%'" {
		fmt.Print(q)
		t.Error()
	}
}

func TestNodeWhenEndsWith_ExpectNoError(t *testing.T) {
	colnode := node.NewColumn("name", NewNodeCompiler())
	endsWith := node.NewEndsWith(NewNodeCompiler())

	endsWith.Left = colnode
	endsWith.Right = node.NewString("ace", NewNodeCompiler())

	s := NewSelectCompiler("u_user", endsWith)
	q, _ := s.Compile()

	if q != "SELECT * FROM u_user WHERE name LIKE '%ace'" {
		fmt.Print(q)
		t.Error()
	}
}
