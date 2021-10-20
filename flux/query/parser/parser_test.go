package parser

import (
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
	encoded := "age = 45"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error()
	}

	sql, _ := testutil_NodeToSQL("sys_user", root)

	if sql != "SELECT * FROM sys_user WHERE 'age' = '45'" {
		t.Error(sql)
	}
}

func TestParser_02(t *testing.T) {
	encoded := "age = 45 and name = ace"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error()
	}

	sql, _ := testutil_NodeToSQL("sys_user", root)

	if sql != "SELECT * FROM sys_user WHERE ( 'age' = '45' AND 'name' = 'ace' )" {
		t.Error(sql)
	}
}

func TestParser_03(t *testing.T) {
	encoded := "age = 45 and name = ace or loser = no"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error()
	}

	sql, _ := testutil_NodeToSQL("sys_user", root)

	if sql != "SELECT * FROM sys_user WHERE ( 'age' = '45' AND ( 'name' = 'ace' OR 'loser' = 'no' ) )" {
		t.Error(sql)
	}
}

func TestParser_04(t *testing.T) {
	encoded := "(loser = no)"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE 'loser' = 'no'" {
		t.Error(sql)
	}
}

func TestParser_05(t *testing.T) {
	encoded := "(((loser = no)))"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE 'loser' = 'no'" {
		t.Error(sql)
	}
}

func TestParser_06(t *testing.T) {
	encoded := "(((loser = no)) or age = 45)"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE ( 'loser' = 'no' OR 'age' = '45' )" {
		t.Error(sql)
	}
}

func TestParser_07(t *testing.T) {
	encoded := "( (a = 1 and b = 2) or d = 3)"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE ( ( 'a' = '1' AND 'b' = '2' ) OR 'd' = '3' )" {
		t.Error(sql)
	}
}

func TestParser_08(t *testing.T) {
	encoded := "( (a = 1 and b = 2) or ((((d = 3)))))"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE ( ( 'a' = '1' AND 'b' = '2' ) OR 'd' = '3' )" {
		t.Error(sql)
	}
}

func TestParser_09(t *testing.T) {
	encoded := "a = 1 and (b = 2 or x = 3)"

	root, err := Parse(encoded, compiler)
	if err != nil {
		t.Error(err)
	}

	sql, err2 := testutil_NodeToSQL("sys_user", root)
	if err2 != nil {
		t.Error(err2)
	}

	if sql != "SELECT * FROM sys_user WHERE ( ( 'a' = '1' AND 'b' = '2' ) OR 'd' = '3' )" {
		t.Error(sql)
	}
}
