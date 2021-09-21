package query

import (
	"fmt"
	"testing"
)

func TestWhenEqual_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	builder.Not()
	_ = builder.Column("first")
	_ = builder.Equals()
	_ = builder.StringValue("last")

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE NOT ( first = 'last' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNotAnd_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	builder.Not()
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

	if sql != "SELECT * FROM u_user WHERE NOT ( ( first = 'last' AND last = 'first' ) )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNotOr_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	builder.Not()
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

	if sql != "SELECT * FROM u_user WHERE NOT ( ( first = 'last' OR last = 'first' ) )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNotAndAnd_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	builder.Not()
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

	if sql != "SELECT * FROM u_user WHERE NOT ( ( ( first = 'last' AND last = 'first' ) AND age = 'old' ) )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNotAndOr_ExpectNoError(t *testing.T) {
	builder := newChainBuilder(newTestUtil_NodeCompiler())

	builder.Not()
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

	if sql != "SELECT * FROM u_user WHERE NOT ( ( ( first = 'last' AND last = 'first' ) OR age = 'old' ) )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNotAndAndAnd_ExpectNoError(t *testing.T) {
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
	builder.Not()

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE NOT ( ( ( ( first = 'last' AND last = 'first' ) AND age = 'old' ) AND one = 'two' ) )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNotAndAndOr_ExpectNoError(t *testing.T) {
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
	builder.Not()

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE NOT ( ( ( ( first = 'last' AND last = 'first' ) AND age = 'old' ) OR one = 'two' ) )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNotAndOrAndOr_ExpectNoError(t *testing.T) {
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
	builder.Not()

	root, _ := builder.Root()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE NOT ( ( ( ( ( first = 'last' AND last = 'first' ) OR age = 'old' ) AND one = 'two' ) OR 3 = '4' ) )" {
		fmt.Println(sql)
		t.Error()
	}
}
