package query

import (
	"fmt"

	"github.com/amortaza/aceql/flux/node"
)

type queryState int

const (
	expectLHS queryState = iota + 1
	expectRHS
	expectCompareOperator
	expectBooleanOperator
)

type chainBuilder struct {
	root                node.Node
	lhs                 node.Node
	rhs                 node.Node
	headCompareOperator node.Node
	tailBooleanOperator node.Node
	tailCompareOperator node.Node
	state               queryState
	not                 bool

	nodeCompiler node.Compiler
}

func newChainBuilder(compiler node.Compiler) *chainBuilder {
	return &chainBuilder{
		state:        expectLHS,
		nodeCompiler: compiler,
	}
}

func (builder *chainBuilder) HasRoot() bool {
	return builder.root != nil
}

func (builder *chainBuilder) Root() (node.Node, error) {
	if builder.root == nil {
		return nil, nil
	}

	if builder.not {
		not := node.NewNot(builder.nodeCompiler)

		err := not.Put(builder.root)
		if err != nil {
			return nil, err
		}

		builder.root = not
	}

	return builder.root, nil
}

func (builder *chainBuilder) Column(name string) error {
	if builder.state != expectLHS {
		return fmt.Errorf("was not expecting LHS")
	}

	builder.lhs = node.NewColumn(name, builder.nodeCompiler)
	builder.state = expectCompareOperator

	return nil
}

func (builder *chainBuilder) StringValue(value string) error {
	if builder.state != expectRHS {
		return fmt.Errorf("was not expecting RHS")
	}

	builder.rhs = node.NewString(value, builder.nodeCompiler)
	builder.state = expectBooleanOperator

	return builder.closeCompare()
}

func (builder *chainBuilder) NumberValue(value float32) error {
	if builder.state != expectRHS {
		return fmt.Errorf("was not expecting RHS")
	}

	builder.rhs = node.NewNumber(value, builder.nodeCompiler)
	builder.state = expectBooleanOperator

	return builder.closeCompare()
}

func (builder *chainBuilder) Equals() error {
	if builder.state != expectCompareOperator {
		return fmt.Errorf("was not expecting Compare operator chainBuilder.Equals()")
	}

	builder.tailCompareOperator = node.NewEquals(builder.nodeCompiler)
	builder.state = expectRHS

	return nil
}

func (builder *chainBuilder) NotEquals() error {
	if builder.state != expectCompareOperator {
		return fmt.Errorf("was not expecting Compare operator chainBuilder.NotEquals()")
	}

	builder.tailCompareOperator = node.NewNotEquals(builder.nodeCompiler)
	builder.state = expectRHS

	return nil
}

func (builder *chainBuilder) LessThan() error {
	if builder.state != expectCompareOperator {
		return fmt.Errorf("was not expecting Compare operator chainBuilder.LessThan()")
	}

	builder.tailCompareOperator = node.NewLessThan(builder.nodeCompiler)
	builder.state = expectRHS

	return nil
}

func (builder *chainBuilder) LessOrEqual() error {
	if builder.state != expectCompareOperator {
		return fmt.Errorf("was not expecting Compare operator chainBuilder.LessThan()")
	}

	builder.tailCompareOperator = node.NewLessOrEquals(builder.nodeCompiler)
	builder.state = expectRHS

	return nil
}

func (builder *chainBuilder) GreaterThan() error {
	if builder.state != expectCompareOperator {
		return fmt.Errorf("was not expecting Compare operator chainBuilder.GreaterThan()")
	}

	builder.tailCompareOperator = node.NewGreaterThan(builder.nodeCompiler)
	builder.state = expectRHS

	return nil
}

func (builder *chainBuilder) GreaterOrEqual() error {
	if builder.state != expectCompareOperator {
		return fmt.Errorf("was not expecting Compare operator chainBuilder.GreaterOrEqual()")
	}

	builder.tailCompareOperator = node.NewGreaterOrEquals(builder.nodeCompiler)
	builder.state = expectRHS

	return nil
}

func (builder *chainBuilder) StartsWith() error {
	if builder.state != expectCompareOperator {
		return fmt.Errorf("was not expecting Compare operator chainBuilder.StartsWith()")
	}

	builder.tailCompareOperator = node.NewStartsWith(builder.nodeCompiler)
	builder.state = expectRHS

	return nil
}

func (builder *chainBuilder) NotStartsWith() error {
	if builder.state != expectCompareOperator {
		return fmt.Errorf("was not expecting Compare operator chainBuilder.NotStartsWith()")
	}

	builder.tailCompareOperator = node.NewNotStartsWith(builder.nodeCompiler)
	builder.state = expectRHS

	return nil
}

func (builder *chainBuilder) EndsWith() error {
	if builder.state != expectCompareOperator {
		return fmt.Errorf("was not expecting Compare operator")
	}

	builder.tailCompareOperator = node.NewEndsWith(builder.nodeCompiler)
	builder.state = expectRHS

	return nil
}

func (builder *chainBuilder) NotEndsWith() error {
	if builder.state != expectCompareOperator {
		return fmt.Errorf("was not expecting Compare operator")
	}

	builder.tailCompareOperator = node.NewNotEndsWith(builder.nodeCompiler)
	builder.state = expectRHS

	return nil
}

func (builder *chainBuilder) Contains() error {
	if builder.state != expectCompareOperator {
		return fmt.Errorf("was not expecting Compare operator")
	}

	builder.tailCompareOperator = node.NewContains(builder.nodeCompiler)
	builder.state = expectRHS

	return nil
}

func (builder *chainBuilder) NotContains() error {
	if builder.state != expectCompareOperator {
		return fmt.Errorf("was not expecting Compare operator")
	}

	builder.tailCompareOperator = node.NewNotContains(builder.nodeCompiler)
	builder.state = expectRHS

	return nil
}

func (builder *chainBuilder) Node(n node.Node) error {
	builder.tailCompareOperator = n
	return builder.closeCompareNode()
}

func (builder *chainBuilder) closeCompareNode() error {
	if builder.headCompareOperator == nil {
		builder.headCompareOperator = builder.tailCompareOperator
		builder.root = builder.headCompareOperator
	}

	if builder.tailBooleanOperator != nil {
		err := builder.tailBooleanOperator.Put(builder.tailCompareOperator)
		if err != nil {
			return err
		}
	}

	builder.tailCompareOperator = nil
	builder.state = expectBooleanOperator

	return nil
}

func (builder *chainBuilder) closeCompare() error {
	err := builder.tailCompareOperator.Put(builder.lhs)
	if err != nil {
		return err
	}

	builder.lhs = nil

	err = builder.tailCompareOperator.Put(builder.rhs)
	if err != nil {
		return err
	}

	builder.rhs = nil

	if builder.headCompareOperator == nil {
		builder.headCompareOperator = builder.tailCompareOperator
		builder.root = builder.headCompareOperator
	}

	if builder.tailBooleanOperator != nil {
		err := builder.tailBooleanOperator.Put(builder.tailCompareOperator)
		if err != nil {
			return err
		}
	}

	builder.tailCompareOperator = nil
	builder.state = expectBooleanOperator

	return nil
}

func (builder *chainBuilder) And() error {
	if builder.state != expectBooleanOperator {
		return fmt.Errorf("was not expecting Boolean operator in chainBuilder.And()")
	}

	and := node.NewAnd(builder.nodeCompiler)

	if builder.tailBooleanOperator != nil {
		err := and.Put(builder.tailBooleanOperator)
		if err != nil {
			return err
		}
	} else {
		err := and.Put(builder.root)
		if err != nil {
			return err
		}
	}

	builder.root = and
	builder.tailBooleanOperator = and

	builder.state = expectLHS

	return nil
}

func (builder *chainBuilder) Or() error {
	if builder.state != expectBooleanOperator {
		return fmt.Errorf("was not expecting Boolean operator in chainBuilder.Or()")
	}

	or := node.NewOr(builder.nodeCompiler)

	if builder.tailBooleanOperator != nil {
		err := or.Put(builder.tailBooleanOperator)
		if err != nil {
			return err
		}
	} else {
		err := or.Put(builder.root)
		if err != nil {
			return err
		}
	}

	builder.root = or
	builder.tailBooleanOperator = or

	builder.state = expectLHS

	return nil
}

func (builder *chainBuilder) Not() {
	builder.not = true
}
