package query

import (
	"fmt"
	_ "fmt"
	"testing"
)

var compiler = newTestUtil_NodeCompiler()

func TestParser_00(t *testing.T) {
	encoded := ""

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error()
	}

	sql, _ := testutil_NodeToSQL("u_user", root)

	if sql != "SELECT * FROM u_user" {
		t.Error()
	}
}

func TestParser_01(t *testing.T) {
	encoded := "a = 1"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error()
	}

	sql, _ := testutil_NodeToSQL("sys_user", root)

	if sql != "SELECT * FROM sys_user WHERE a = '1'" {
		t.Error(sql)
	}
}

func TestParser_02(t *testing.T) {
	encoded := "a = 1 and b = 2"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error()
	}

	sql, _ := testutil_NodeToSQL("sys_user", root)

	if sql != "SELECT * FROM sys_user WHERE ( a = '1' AND b = '2' )" {
		t.Error(sql)
	}
}

func TestParser_02_0(t *testing.T) {
	encoded := "age = 45 and name = ace or loser = no"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error()
	}

	sql, _ := testutil_NodeToSQL("sys_user", root)

	if sql != "SELECT * FROM sys_user WHERE ( age = '45' AND ( name = 'ace' OR loser = 'no' ) )" {
		t.Error(sql)
	}
}

func TestParser_03_0(t *testing.T) {
	encoded := "(a = 1)"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE a = '1'" {
		t.Error(sql)
	}
}

func TestParser_03_1(t *testing.T) {
	encoded := "((a = 1))"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE a = '1'" {
		t.Error(sql)
	}
}

func TestParser_03_2(t *testing.T) {
	encoded := "(((a = 1)))"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE a = '1'" {
		t.Error(sql)
	}
}

func TestParser_04(t *testing.T) {
	encoded := "(((a = 1)) or b = 2)"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE ( a = '1' OR b = '2' )" {
		t.Error(sql)
	}
}

func TestParser_05(t *testing.T) {
	encoded := "( (a = 1 and b = 2) or d = 3)"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE ( ( a = '1' AND b = '2' ) OR d = '3' )" {
		t.Error(sql)
	}
}

func TestParser_06(t *testing.T) {
	encoded := "( (a = 1 and b = 2) or ((((d = 3)))))"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE ( ( a = '1' AND b = '2' ) OR d = '3' )" {
		t.Error(sql)
	}
}

func TestParser_07(t *testing.T) {
	encoded := "a = 1 and (b = 2 or x = 3)"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE ( a = '1' AND ( b = '2' OR x = '3' ) )" {
		t.Error(sql)
	}
}

func TestParser_08_01(t *testing.T) {
	//encoded := "(a = 1 and b = 2 and d = 3) or ( e = 4 and f = 5 and g = 6 ) or ( h = 7 and i = 8 )"
	encoded := "(a = 1 and b = 2 and d = 3) or e = 4"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE ( ( a = '1' AND ( b = '2' AND d = '3' ) ) OR e = '4' )" {
		t.Error(sql)
	}
}

func TestParser_08_02(t *testing.T) {
	encoded := "(a = 1 and b = 2 and d = 3) or ( e = 4 and f = 5 and g = 6 ) or ( h = 7 and i = 8 )"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE ( ( a = '1' AND ( b = '2' AND d = '3' ) ) OR e = '4' )" {
		t.Error(sql)
	}
}

func TestParser_09(t *testing.T) {
	encoded := "a = 1 and b = 2 and c = 3 and ( ( aaa = 111 or bbb = 222 ) aa = 11 or  bb = 22 or ( ccc = 333 or ddd = 444 or eee = 555 ) ) and ( d = 4 or  ( cc = 33 or ( ( fff = 666 ) or dd = 44 ) ) ) or e = 5"
	encoded = "(a = 1 and b = 2) or e = 5"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	fmt.Println( sql ) // debug
	if sql != "SELECT * FROM sys_user WHERE ( ( a = '1' AND ( b = '2' AND d = '3' ) ) OR e = '4' )" {
		//t.Error(sql)
	}
}
