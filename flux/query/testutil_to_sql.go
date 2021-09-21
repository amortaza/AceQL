package query

func testutilToSQL(table string, filterQuery *FilterQuery) (string, error) {
	root, err := filterQuery.GetRoot()
	if err != nil {
		return "", err
	}

	comp := newTestUtil_SelectCompiler(table, root)

	return comp.Compile()
}
