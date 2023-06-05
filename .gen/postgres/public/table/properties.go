//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Properties = newPropertiesTable("public", "properties", "")

type propertiesTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	Code      postgres.ColumnString
	Name      postgres.ColumnString
	ProductID postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PropertiesTable struct {
	propertiesTable

	EXCLUDED propertiesTable
}

// AS creates new PropertiesTable with assigned alias
func (a PropertiesTable) AS(alias string) *PropertiesTable {
	return newPropertiesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PropertiesTable with assigned schema name
func (a PropertiesTable) FromSchema(schemaName string) *PropertiesTable {
	return newPropertiesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PropertiesTable with assigned table prefix
func (a PropertiesTable) WithPrefix(prefix string) *PropertiesTable {
	return newPropertiesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PropertiesTable with assigned table suffix
func (a PropertiesTable) WithSuffix(suffix string) *PropertiesTable {
	return newPropertiesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPropertiesTable(schemaName, tableName, alias string) *PropertiesTable {
	return &PropertiesTable{
		propertiesTable: newPropertiesTableImpl(schemaName, tableName, alias),
		EXCLUDED:        newPropertiesTableImpl("", "excluded", ""),
	}
}

func newPropertiesTableImpl(schemaName, tableName, alias string) propertiesTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		CodeColumn      = postgres.StringColumn("code")
		NameColumn      = postgres.StringColumn("name")
		ProductIDColumn = postgres.IntegerColumn("product_id")
		allColumns      = postgres.ColumnList{IDColumn, CodeColumn, NameColumn, ProductIDColumn}
		mutableColumns  = postgres.ColumnList{CodeColumn, NameColumn, ProductIDColumn}
	)

	return propertiesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Code:      CodeColumn,
		Name:      NameColumn,
		ProductID: ProductIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
