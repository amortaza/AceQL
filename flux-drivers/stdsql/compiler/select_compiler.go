package compiler

import (
	"github.com/amortaza/aceql/flux/node"
)

type SelectCompiler struct {
	From  string
	Where node.Node
}

func NewSelectCompiler(table string, where node.Node) *SelectCompiler {
	s := &SelectCompiler{}

	s.From = table
	s.Where = where

	return s
}

func (s *SelectCompiler) Compile() (string, error) {
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
