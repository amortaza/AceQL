package query

import (
	"fmt"
	"testing"
)

func TestWhenCompareToNumber_ExpectNoError(t *testing.T) {
	filterQuery := NewFilterQuery(newTestUtil_NodeCompiler())

	_ = filterQuery.AddNumber("u_name", Equals, 1)

	sql, _ := testutil_FilterQueryToSQL("u_user", filterQuery)

	if sql != "SELECT * FROM u_user WHERE u_name = 1.000000" {
		fmt.Println(sql)
		t.Error()
	}
}
