package compiler

import (
	"github.com/amortaza/aceql/flux/node"
	"strconv"
	"strings"
)

type SelectCompiler struct {
	From  string
	Columns []string
	Where node.Node
}

func NewSelectCompiler(table string, columns []string, where node.Node) *SelectCompiler {
	s := &SelectCompiler{}

	s.From = table
	s.Columns = columns
	s.Where = where

	return s
}

func (s *SelectCompiler) Compile(paginationIndex int, paginationSize int) (string, string, error) {
	q := "SELECT " + strings.Join( s.Columns[:],", ") + " FROM " + s.From
	queryUsedForCount := "SELECT COUNT(1) AS total FROM " + s.From

	if s.Where != nil {
		sql, err := s.Where.Compile()
		if err != nil {
			return "", "", err
		}

		if sql != "" {
			q += " WHERE " + sql
			queryUsedForCount += " WHERE " + sql
		}
	}

	if paginationSize > -1 {
		q += " LIMIT " + strconv.Itoa(paginationIndex) + ", " + strconv.Itoa(paginationSize)
	}

	return q, queryUsedForCount, nil
}
