package query

import (
	"fmt"
	"github.com/amortaza/aceql/flux/node"
)

type FilterQuery struct {
	b1 *chainBuilder
	b2 *chainBuilder

	encodedQuery string

	compiler node.Compiler
}

func NewFilterQuery(compiler node.Compiler) *FilterQuery {
	return &FilterQuery{
		compiler: compiler,
		b1:       newChainBuilder(compiler),
	}
}

func (query *FilterQuery) SetEncodedQuery(encodedQuery string) {
	query.encodedQuery = encodedQuery
}

func (query *FilterQuery) Not() {
	query.b1.Not()
}

func (query *FilterQuery) Add(field string, op OpType, rhs string) error {
	var err error

	if query.b1.HasRoot() {
		err = query.b1.And()
		if err != nil {
			return err
		}
	}

	err = query.b1.Column(field)
	if err != nil {
		return err
	}

	err = query.compareStringOp(op)
	if err != nil {
		return err
	}

	err = query.b1.StringValue(rhs)
	if err != nil {
		return err
	}

	return nil
}

func (query *FilterQuery) AddNumber(field string, op OpType, rhs float32) error {
	var err error

	if query.b1.HasRoot() {
		err = query.b1.And()
		if err != nil {
			return err
		}
	}

	err = query.b1.Column(field)
	if err != nil {
		return err
	}

	err = query.compareNumberOp(op)
	if err != nil {
		return err
	}

	err = query.b1.NumberValue(rhs)
	if err != nil {
		return err
	}

	return nil
}

func (query *FilterQuery) compareStringOp(op OpType) error {
	var err error

	if op == Equals {
		err = query.b1.Equals()
	} else if op == NotEquals {
		err = query.b1.NotEquals()
	} else if op == LessThan {
		err = query.b1.LessThan()
	} else if op == LessOrEqual {
		err = query.b1.LessOrEqual()
	} else if op == GreaterThan {
		err = query.b1.GreaterThan()
	} else if op == GreaterOrEqual {
		err = query.b1.GreaterOrEqual()
	} else if op == StartsWith {
		err = query.b1.StartsWith()
	} else if op == NotStartsWith {
		err = query.b1.NotStartsWith()
	} else if op == EndsWith {
		err = query.b1.EndsWith()
	} else if op == NotEndsWith {
		err = query.b1.NotEndsWith()
	} else if op == Contains {
		err = query.b1.Contains()
	} else if op == NotContains {
		err = query.b1.NotContains()
	} else {
		err = fmt.Errorf("filterQuery.compareOp does not recognize OpType")
	}

	return err
}

func (query *FilterQuery) compareNumberOp(op OpType) error {
	var err error

	if op == Equals {
		err = query.b1.Equals()
	} else if op == NotEquals {
		err = query.b1.NotEquals()
	} else if op == LessThan {
		err = query.b1.LessThan()
	} else if op == LessOrEqual {
		err = query.b1.LessOrEqual()
	} else if op == GreaterThan {
		err = query.b1.GreaterThan()
	} else if op == GreaterOrEqual {
		err = query.b1.GreaterOrEqual()
	} else {
		err = fmt.Errorf("filterQuery.compareOp does not recognize OpType")
	}

	return err
}

func (query *FilterQuery) AddOr(field string, op OpType, rhs string) error {
	var err error

	if query.b1.HasRoot() {
		err = query.b1.Or()
		if err != nil {
			return err
		}
	}

	err = query.b1.Column(field)
	if err != nil {
		return err
	}

	err = query.compareStringOp(op)
	if err != nil {
		return err
	}

	err = query.b1.StringValue(rhs)
	if err != nil {
		return err
	}

	return nil
}

func (query *FilterQuery) AddOrNumber(field string, op OpType, rhs float32) error {
	var err error

	if query.b1.HasRoot() {
		err = query.b1.Or()
		if err != nil {
			return err
		}
	}

	err = query.b1.Column(field)
	if err != nil {
		return err
	}

	err = query.compareNumberOp(op)
	if err != nil {
		return err
	}

	err = query.b1.NumberValue(rhs)
	if err != nil {
		return err
	}

	return nil
}

func (query *FilterQuery) AndGroup() error {
	if query.b2 == nil {
		query.b2 = newChainBuilder(query.compiler)
	}

	root, err := query.b1.Root()
	if err != nil {
		return err
	}

	err = query.b2.Node(root)
	if err != nil {
		return err
	}

	err = query.b2.And()
	if err != nil {
		return err
	}

	query.b1 = newChainBuilder(query.compiler)

	return nil
}

func (query *FilterQuery) OrGroup() error {
	var err error

	if query.b2 == nil {
		query.b2 = newChainBuilder(query.compiler)
	}

	root, err2 := query.b1.Root()
	if err2 != nil {
		return err2
	}

	err = query.b2.Node(root)
	if err != nil {
		return err
	}

	err = query.b2.Or()
	if err != nil {
		return err
	}

	query.b1 = newChainBuilder(query.compiler)

	return nil
}

func (query *FilterQuery) GetRoot() (node.Node, error) {
	if query.encodedQuery == "" {
		return query.getRoot()
	}

	return Parse( query.encodedQuery, query.compiler )
}

func (query *FilterQuery) getRoot() (node.Node, error) {
	var root node.Node
	var err error

	root, err = query.b1.Root()
	if err != nil {
		return nil, err
	}

	if query.b2 != nil {
		err = query.b2.Node(root)
		if err != nil {
			return nil, err
		}

		root, err = query.b2.Root()
		if err != nil {
			return nil, err
		}
	}

	return root, nil
}
