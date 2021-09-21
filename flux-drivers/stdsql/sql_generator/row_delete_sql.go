package sql_generator

import (
	"fmt"
)

type RowDelete_SqlGenerator struct {}

func NewRowDelete_SqlGenerator() *RowDelete_SqlGenerator {
	return &RowDelete_SqlGenerator{}
}

func (generator *RowDelete_SqlGenerator) GenerateDeleteSQL(table string, pk string) string {
	return fmt.Sprintf("DELETE FROM `%s` WHERE x_id = '%s';", table, pk)
}
