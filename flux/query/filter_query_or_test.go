package query

import (
	"fmt"
	"testing"
)

func TestQueryWhenOr_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", Equals, "u_id")
	_ = filterQuery.AddOr("u_age", Equals, "u_years")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( u_name = 'u_id' OR u_age = 'u_years' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenOrOr_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", Equals, "u_id")
	_ = filterQuery.AddOr("u_age", Equals, "u_years")
	_ = filterQuery.AddOr("u_date", Equals, "u_today")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( u_name = 'u_id' OR u_age = 'u_years' ) OR u_date = 'u_today' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestQueryWhenAndOr_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", Equals, "u_id")
	_ = filterQuery.Add("u_age", Equals, "u_years")
	_ = filterQuery.AddOr("u_date", Equals, "u_today")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( u_name = 'u_id' AND u_age = 'u_years' ) OR u_date = 'u_today' )" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenOrGroup_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", Equals, "u_id")
	_ = filterQuery.OrGroup()
	_ = filterQuery.Add("u_age", Equals, "u_years")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( u_name = 'u_id' OR u_age = 'u_years' )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}

func TestWhenAndOrGroupAnd_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("a", Equals, "b")
	_ = filterQuery.Add("c", Equals, "d")
	_ = filterQuery.OrGroup()
	_ = filterQuery.Add("e", Equals, "f")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( a = 'b' AND c = 'd' ) OR e = 'f' )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}

func TestWhenAndAndOrGroupAndAnd_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("a", Equals, "b")
	_ = filterQuery.Add("c", Equals, "d")
	_ = filterQuery.OrGroup()
	_ = filterQuery.Add("e", Equals, "f")
	_ = filterQuery.Add("g", Equals, "h")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( a = 'b' AND c = 'd' ) OR ( e = 'f' AND g = 'h' ) )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}

func TestWhenComplexOrGroup_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("a", Equals, "b")
	_ = filterQuery.OrGroup()
	_ = filterQuery.Add("c", Equals, "d")
	_ = filterQuery.Add("e", Equals, "f")
	_ = filterQuery.OrGroup()
	_ = filterQuery.Add("g", Equals, "h")
	_ = filterQuery.Add("i", Equals, "j")
	_ = filterQuery.Add("k", Equals, "l")
	_ = filterQuery.OrGroup()
	_ = filterQuery.Add("m", Equals, "n")
	_ = filterQuery.Add("o", Equals, "p")
	_ = filterQuery.Add("q", Equals, "r")
	_ = filterQuery.Add("s", Equals, "t")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE ( ( ( a = 'b' OR ( c = 'd' AND e = 'f' ) ) OR ( ( g = 'h' AND i = 'j' ) AND k = 'l' ) ) OR ( ( ( m = 'n' AND o = 'p' ) AND q = 'r' ) AND s = 't' ) )" {
		fmt.Printf("'%s'", sql)
		t.Error()
	}
}
