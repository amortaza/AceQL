package query

func testutil_FilterQueryToSQL(table string, filterQuery *FilterQuery) (string, error) {
	root, err := filterQuery.GetRoot()
	if err != nil {
		return "", err
	}

	comp := newTestUtil_SelectCompiler(table, root)

	return comp.Compile()
}
