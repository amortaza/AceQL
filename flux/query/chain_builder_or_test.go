package query

import (
	"fmt"
	"testing"
)

func TestWhenOr_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	_ = builder.Column("first")
	_ = builder.Equals()
	_ = builder.StringValue("last")
	_ = builder.Or()
	_ = builder.Column("last")
	_ = builder.Equals()
	_ = builder.StringValue("first")

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE ( first = 'last' OR last = 'first' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenAndOr_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	_ = builder.Column("first")
	_ = builder.Equals()
	_ = builder.StringValue("last")
	_ = builder.And()
	_ = builder.Column("last")
	_ = builder.Equals()
	_ = builder.StringValue("first")
	_ = builder.Or()
	_ = builder.Column("age")
	_ = builder.Equals()
	_ = builder.StringValue("old")

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE ( ( first = 'last' AND last = 'first' ) OR age = 'old' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenAndAndOr_ExpectNoError(t *testing.T) {
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
	_ = builder.Or()
	_ = builder.Column("one")
	_ = builder.Equals()
	_ = builder.StringValue("two")

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE ( ( ( first = 'last' AND last = 'first' ) AND age = 'old' ) OR one = 'two' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenAndOrAndOr_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	_ = builder.Column("first")
	_ = builder.Equals()
	_ = builder.StringValue("last")
	_ = builder.And()
	_ = builder.Column("last")
	_ = builder.Equals()
	_ = builder.StringValue("first")
	_ = builder.Or()
	_ = builder.Column("age")
	_ = builder.Equals()
	_ = builder.StringValue("old")
	_ = builder.And()
	_ = builder.Column("one")
	_ = builder.Equals()
	_ = builder.StringValue("two")
	_ = builder.Or()
	_ = builder.Column("3")
	_ = builder.Equals()
	_ = builder.StringValue("4")

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE ( ( ( ( first = 'last' AND last = 'first' ) OR age = 'old' ) AND one = 'two' ) OR 3 = '4' )" {
		fmt.Println(sql)
		t.Error()
	}
}
