package parser

import (
	"testing"
)

func TestParserWhenEmpty_ExpectNoError(t *testing.T) {
	encoded := ""

	root, err := Parse(encoded)
	if err != nil {
		t.Error()
	}

	sql, _ := testutil_NodeToSQL("u_user", root)

	if sql != "SELECT * FROM u_user" {
		t.Error()
	}
}

func TestParserWhenEqualsNumber_ExpectNoError(t *testing.T) {
	encoded := "age = 45"

	root, err := Parse(encoded)
	if err != nil {
		t.Error()
	}

	sql, _ := testutil_NodeToSQL("u_user", root)

	if sql != "SELECT * FROM u_user where age = 45" {
		t.Error()
	}
}
/*
func TestWhenEqualsString_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", Equals, "u_id")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE u_name = 'u_id'" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenAnd_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", Equals, "u_id")
	_ = filterQuery.Add("u_age", Equals, "u_years")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( u_name = 'u_id' AND u_age = 'u_years' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenAndAnd_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", Equals, "u_id")
	_ = filterQuery.Add("u_age", Equals, "u_years")
	_ = filterQuery.Add("u_date", Equals, "u_today")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( u_name = 'u_id' AND u_age = 'u_years' ) AND u_date = 'u_today' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenAndGroup_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", Equals, "u_id")
	_ = filterQuery.AndGroup()
	_ = filterQuery.Add("u_age", Equals, "u_years")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( u_name = 'u_id' AND u_age = 'u_years' )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}

func TestWhenAndAndGroupAnd_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("a", Equals, "b")
	_ = filterQuery.Add("c", Equals, "d")
	_ = filterQuery.AndGroup()
	_ = filterQuery.Add("e", Equals, "f")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( a = 'b' AND c = 'd' ) AND e = 'f' )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}

func TestWhenAndAndGroupAndAnd_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("a", Equals, "b")
	_ = filterQuery.Add("c", Equals, "d")
	_ = filterQuery.AndGroup()
	_ = filterQuery.Add("e", Equals, "f")
	_ = filterQuery.Add("g", Equals, "h")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( a = 'b' AND c = 'd' ) AND ( e = 'f' AND g = 'h' ) )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}

func TestWhenComplexAndGroup_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("a", Equals, "b")
	_ = filterQuery.AndGroup()
	_ = filterQuery.Add("c", Equals, "d")
	_ = filterQuery.Add("e", Equals, "f")
	_ = filterQuery.AndGroup()
	_ = filterQuery.Add("g", Equals, "h")
	_ = filterQuery.Add("i", Equals, "j")
	_ = filterQuery.Add("k", Equals, "l")
	_ = filterQuery.AndGroup()
	_ = filterQuery.Add("m", Equals, "n")
	_ = filterQuery.Add("o", Equals, "p")
	_ = filterQuery.Add("q", Equals, "r")
	_ = filterQuery.Add("s", Equals, "t")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( ( a = 'b' AND ( c = 'd' AND e = 'f' ) ) AND ( ( g = 'h' AND i = 'j' ) AND k = 'l' ) ) AND ( ( ( m = 'n' AND o = 'p' ) AND q = 'r' ) AND s = 't' ) )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}
*/