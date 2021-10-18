package query

import (
	"fmt"
	"testing"
)

func TestWhenNotEqual_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	filterQuery.Not()
	_ = filterQuery.Add("u_name", Equals, "u_id")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE NOT ( u_name = 'u_id' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestQueryWhenNotAnd_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	filterQuery.Not()
	_ = filterQuery.Add("u_name", Equals, "u_id")
	_ = filterQuery.Add("u_age", Equals, "u_years")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE NOT ( ( u_name = 'u_id' AND u_age = 'u_years' ) )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestQueryWhenNotAndAnd_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	filterQuery.Not()
	_ = filterQuery.Add("u_name", Equals, "u_id")
	_ = filterQuery.Add("u_age", Equals, "u_years")
	_ = filterQuery.Add("u_date", Equals, "u_today")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE NOT ( ( ( u_name = 'u_id' AND u_age = 'u_years' ) AND u_date = 'u_today' ) )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNotAndGroupedWithAnd_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	filterQuery.Not()
	_ = filterQuery.Add("u_name", Equals, "u_id")
	_ = filterQuery.AndGroup()
	_ = filterQuery.Add("u_age", Equals, "u_years")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( NOT ( u_name = 'u_id' ) AND u_age = 'u_years' )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}

func TestWhenAndNot_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", Equals, "u_id")
	_ = filterQuery.AndGroup()
	filterQuery.Not()
	_ = filterQuery.Add("u_age", Equals, "u_years")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( u_name = 'u_id' AND NOT ( u_age = 'u_years' ) )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}

func TestWhenNot2AndsWithSeparateAnd_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	filterQuery.Not()
	_ = filterQuery.Add("a", Equals, "b")
	_ = filterQuery.Add("c", Equals, "d")
	_ = filterQuery.AndGroup()
	_ = filterQuery.Add("e", Equals, "f")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( NOT ( ( a = 'b' AND c = 'd' ) ) AND e = 'f' )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}

func TestWhenAndNotAnd_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("a", Equals, "b")
	_ = filterQuery.Add("c", Equals, "d")
	_ = filterQuery.AndGroup()
	filterQuery.Not()
	_ = filterQuery.Add("e", Equals, "f")
	_ = filterQuery.Add("g", Equals, "h")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( a = 'b' AND c = 'd' ) AND NOT ( ( e = 'f' AND g = 'h' ) ) )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}

func TestWhenComplexGroups_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("a", Equals, "b")
	_ = filterQuery.AndGroup()
	_ = filterQuery.Add("c", Equals, "d")
	_ = filterQuery.Add("e", Equals, "f")
	_ = filterQuery.AndGroup()
	filterQuery.Not()
	_ = filterQuery.Add("g", Equals, "h")
	_ = filterQuery.Add("i", Equals, "j")
	_ = filterQuery.Add("k", Equals, "l")
	_ = filterQuery.AndGroup()
	_ = filterQuery.Add("m", Equals, "n")
	_ = filterQuery.Add("o", Equals, "p")
	_ = filterQuery.Add("q", Equals, "r")
	_ = filterQuery.Add("s", Equals, "t")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( ( a = 'b' AND ( c = 'd' AND e = 'f' ) ) AND NOT ( ( ( g = 'h' AND i = 'j' ) AND k = 'l' ) ) ) AND ( ( ( m = 'n' AND o = 'p' ) AND q = 'r' ) AND s = 't' ) )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}

func TestWhenComplexGroupsVariation_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("a", Equals, "b")
	_ = filterQuery.AndGroup()
	_ = filterQuery.Add("c", Equals, "d")
	_ = filterQuery.Add("e", Equals, "f")
	_ = filterQuery.AndGroup()
	filterQuery.Not()
	_ = filterQuery.Add("g", Equals, "h")
	_ = filterQuery.Add("i", Equals, "j")
	_ = filterQuery.Add("k", Equals, "l")
	_ = filterQuery.AndGroup()
	filterQuery.Not()
	_ = filterQuery.Add("m", Equals, "n")
	_ = filterQuery.Add("o", Equals, "p")
	_ = filterQuery.Add("q", Equals, "r")
	_ = filterQuery.Add("s", Equals, "t")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( ( a = 'b' AND ( c = 'd' AND e = 'f' ) ) AND NOT ( ( ( g = 'h' AND i = 'j' ) AND k = 'l' ) ) ) AND NOT ( ( ( ( m = 'n' AND o = 'p' ) AND q = 'r' ) AND s = 't' ) ) )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}

func TestWhenComplexGroupsAnotherVariation_ExpectNoError(t *testing.T) {
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
	filterQuery.Not()
	_ = filterQuery.Add("m", Equals, "n")
	_ = filterQuery.Add("o", Equals, "p")
	_ = filterQuery.Add("q", Equals, "r")
	_ = filterQuery.Add("s", Equals, "t")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( ( a = 'b' AND ( c = 'd' AND e = 'f' ) ) AND ( ( g = 'h' AND i = 'j' ) AND k = 'l' ) ) AND NOT ( ( ( ( m = 'n' AND o = 'p' ) AND q = 'r' ) AND s = 't' ) ) )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}
