package sql_generator

import (
	"fmt"
)

type TableCreator struct {}

func NewTableCreate_SqlGenerator() *TableCreator {
	return &TableCreator{}
}

func (generator *TableCreator) GenerateCreateTableSQL(table string) string {
	return fmt.Sprintf("CREATE TABLE `%s` (`x_id` CHAR(32) NOT NULL, PRIMARY KEY (`x_id`));", table )
}
