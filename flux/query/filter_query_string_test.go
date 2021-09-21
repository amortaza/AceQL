package query

import (
	"fmt"
	"testing"
)

func TestWhenStartsWith_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", StartsWith, "foo")

	root, _ := filterQuery.GetRoot()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE u_name LIKE 'foo%'" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNotStartsWith_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", NotStartsWith, "foo")

	root, _ := filterQuery.GetRoot()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE u_name NOT LIKE 'foo%'" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenEndsWith_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", EndsWith, "foo")

	root, _ := filterQuery.GetRoot()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE u_name LIKE '%foo'" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNotEndsWith_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", NotEndsWith, "foo")

	root, _ := filterQuery.GetRoot()
	comp := newTestUtil_SelectCompiler("u_user", root)

	sql, _ := comp.Compile()

	if sql != "SELECT * FROM u_user WHERE u_name NOT LIKE '%foo'" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenContains_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", Contains, "foo")

	sql, _ := testutilToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE u_name LIKE '%foo%'" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNotContains_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", NotContains, "foo")

	sql, _ := testutilToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE u_name NOT LIKE '%foo%'" {
		fmt.Println(sql)
		t.Error()
	}
}
