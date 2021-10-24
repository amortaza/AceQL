package query

import (
	"fmt"
	"testing"
)

func TestWhenEquals_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", Equals, "u_id")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE u_name = 'u_id'" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenNotEquals_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", NotEquals, "u_id")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE u_name != 'u_id'" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenLessThan_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", LessThan, "u_id")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE u_name < 'u_id'" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenLessThanEqualTo_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", LessOrEquals, "u_id")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE u_name <= 'u_id'" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenGreaterThan_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", GreaterThan, "u_id")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE u_name > 'u_id'" {
		fmt.Println(sql)
		t.Error()
	}
}

func TestWhenGreaterOrEquals_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.Add("u_name", GreaterOrEquals, "u_id")

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE u_name >= 'u_id'" {
		fmt.Println(sql)
		t.Error()
	}
}
