package sql_generator

import (
	"fmt"
)

type TableDelete_SqlGenerator struct {}

func NewTableDelete_SqlGenerator() *TableDelete_SqlGenerator {
	return &TableDelete_SqlGenerator{}
}

func (generator *TableDelete_SqlGenerator) GenerateDeleteTableSQL(table string) string {
	return fmt.Sprintf("DROP TABLE `%s`;", table)
}
