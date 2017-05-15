package database_Celica

import (
	"Celica/checkError_Celica"
)

func (this *CelicaSql) TableCreate(tableName string,
	fieldNames []string, fieldTypes []string) int {
	if this == nil || this.db == nil {
		return 1
	}
	if tableName == "" {
		return 2
	}
	length := len(fieldNames)
	if length < 1 || length != len(fieldTypes) {
		return 3
	}

	tableName = "`" + tableName + "`"
	commad := "create table if not exists " + tableName +
		"(" + fieldNames[0] + " " + fieldTypes[0] + " not null primary key"
	for i := 1; i < length; i++ {
		commad += "," + fieldNames[i] + " " + fieldTypes[i] + " not null"
	}
	commad += ")"

	stmt, err := this.db.Prepare(commad)
	if checkError_Celica.CheckError(err, "TableCreate") == false {
		return -1
	}
	defer stmt.Close()

	stmt.Exec()
	return 0
}

func (this *CelicaSql) TableDrop(tableName string) int {
	if this == nil || this.db == nil {
		return 1
	}
	if tableName == "" {
		return 2
	}

	tableName = "`" + tableName + "`"

	commad := "drop table " + tableName

	stmt, err := this.db.Prepare(commad)
	if checkError_Celica.CheckError(err, "TableDrop") == false {
		return -1
	}
	defer stmt.Close()

	stmt.Exec()
	return 0
}
