package sql_generator

import (
	"fmt"
)

type FieldDelete_SqlGenerator struct {}

func NewFieldDelete_SqlGenerator() *FieldDelete_SqlGenerator {
	return &FieldDelete_SqlGenerator{}
}

func (generator *FieldDelete_SqlGenerator) GenerateDeleteFieldSQL(table string, fieldname string) string {
	sql := fmt.Sprintf("ALTER TABLE `%s` DROP `%s`;", table, fieldname)

	return sql
}
