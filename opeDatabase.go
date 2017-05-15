package database_Celica

import (
	"Celica/checkError_Celica"
)

func (this *CelicaSql) DatabaseCreate(databaseName string) int {
	if this == nil || this.db == nil {
		return 1
	}

	commad := "create database " + databaseName

	stmt, err := this.db.Prepare(commad)
	if checkError_Celica.CheckError(err, "TableDrop") == false {
		return -1
	}
	defer stmt.Close()

	stmt.Exec()
	return 0
}

func (this *CelicaSql) DatabaseDrop(databaseName string) int {
	if this == nil || this.db == nil {
		return 1
	}

	commad := "drop database " + databaseName

	stmt, err := this.db.Prepare(commad)
	if checkError_Celica.CheckError(err, "TableDrop") == false {
		return -1
	}
	defer stmt.Close()

	stmt.Exec()
	return 0
}
