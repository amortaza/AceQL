package query

import (
	"fmt"
	"testing"
)

func TestWhenEmpty_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenSimpleEquals_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	_ = builder.Column("first")
	_ = builder.Equals()
	_ = builder.StringValue("last")

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE first = 'last'" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenAnded_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	_ = builder.Column("first")
	_ = builder.Equals()
	_ = builder.StringValue("last")
	_ = builder.And()
	_ = builder.Column("last")
	_ = builder.Equals()
	_ = builder.StringValue("first")

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE ( first = 'last' AND last = 'first' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNestedAndsWithLeftSidePrioritized_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	_ = builder.Column("first")
	_ = builder.Equals()
	_ = builder.StringValue("last")
	_ = builder.And()
	_ = builder.Column("last")
	_ = builder.Equals()
	_ = builder.StringValue("first")
	_ = builder.And()
	_ = builder.Column("age")
	_ = builder.Equals()
	_ = builder.StringValue("old")

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE ( ( first = 'last' AND last = 'first' ) AND age = 'old' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenThreeAnds_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	_ = builder.Column("first")
	_ = builder.Equals()
	_ = builder.StringValue("last")
	_ = builder.And()
	_ = builder.Column("last")
	_ = builder.Equals()
	_ = builder.StringValue("first")
	_ = builder.And()
	_ = builder.Column("age")
	_ = builder.Equals()
	_ = builder.StringValue("old")
	_ = builder.And()
	_ = builder.Column("one")
	_ = builder.Equals()
	_ = builder.StringValue("two")

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE ( ( ( first = 'last' AND last = 'first' ) AND age = 'old' ) AND one = 'two' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenFourAndsIncludingNumber_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	_ = builder.Column("first")
	_ = builder.Equals()
	_ = builder.StringValue("last")
	_ = builder.And()
	_ = builder.Column("last")
	_ = builder.Equals()
	_ = builder.StringValue("first")
	_ = builder.And()
	_ = builder.Column("age")
	_ = builder.Equals()
	_ = builder.StringValue("old")
	_ = builder.And()
	_ = builder.Column("one")
	_ = builder.Equals()
	_ = builder.StringValue("two")
	_ = builder.And()
	_ = builder.Column("3")
	_ = builder.Equals()
	_ = builder.NumberValue(4)

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE ( ( ( ( first = 'last' AND last = 'first' ) AND age = 'old' ) AND one = 'two' ) AND 3 = 4.000000 )" {
		fmt.Println(sql)
		t.Error()
	}
}
