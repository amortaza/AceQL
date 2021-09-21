package query

import "github.com/amortaza/aceql/flux/node"

type testUtil_SelectCompiler struct {
	From  string
	Where node.Node
}

func newTestUtil_SelectCompiler(table string, where node.Node) *testUtil_SelectCompiler {
	s := &testUtil_SelectCompiler{}

	s.From = table
	s.Where = where

	return s
}

func (s *testUtil_SelectCompiler) Compile() (string, error) {
	q := "SELECT * FROM " + s.From

	if s.Where == nil {
		return q, nil
	}

	sql, err := s.Where.Compile()
	if err != nil {
		return "", err
	}

	if sql != "" {
		q += " WHERE " + sql
	}

	return q, nil
}
