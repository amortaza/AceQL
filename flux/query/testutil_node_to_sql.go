package query

import "github.com/amortaza/aceql/flux/node"

func testutil_NodeToSQL(table string, root node.Node) (string, error) {
	comp := newTestUtil_SelectCompiler(table, root)

	return comp.Compile()
}

