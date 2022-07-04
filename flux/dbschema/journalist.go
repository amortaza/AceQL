package dbschema

type Journalist interface {
	CreateTable(tableName string, tableLabel string) error
	DeleteTable(tableName string) error

	CreateField(tableName string, field *Field) error
	DeleteField(tableName string, fieldname string) error
}
