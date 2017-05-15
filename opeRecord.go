package database_Celica

import (
	"Celica/checkError_Celica"
)

func (this *CelicaSql) RecordInsert(newInsert *RecordOpe) int {
	if this == nil || this.db == nil {
		return 1
	}
	if newInsert.tableName == "" || newInsert.keyName == "" {
		return 2
	}

	commad := "insert into " + newInsert.tableName + "(" + newInsert.keyName + ")value" + "(?)"

	stmt, err := this.db.Prepare(commad)
	if checkError_Celica.CheckError(err, "insert Record") == false {
		return -1
	}
	defer stmt.Close()

	reseult, err := stmt.Exec(newInsert.keyValue)
	if checkError_Celica.CheckError(err, "insert Record") == false {
		return -1
	}

	_, err = reseult.LastInsertId()
	if checkError_Celica.CheckError(err, "insert Record") == false {
		return -1
	}

	return 0
}

func (this *CelicaSql) RecordDelete(nowDel *RecordOpe) int {
	if this == nil || this.db == nil {
		return 1
	}
	if nowDel.tableName == "" || nowDel.keyName == "" {
		return 2
	}

	commad := "delete from " + nowDel.tableName + " where " + nowDel.keyName + " =?"

	stmt, err := this.db.Prepare(commad)
	if checkError_Celica.CheckError(err, "Delete Record") == false {
		return -1
	}
	defer stmt.Close()

	result, err := stmt.Exec(nowDel.keyValue)
	if checkError_Celica.CheckError(err, "Delete Record") == false {
		return -1
	}

	_, err = result.RowsAffected()
	if checkError_Celica.CheckError(err, "Delete Record") == false {
		return -1
	}

	return 0
}

func (this *CelicaSql) RecordCount(nowDel *RecordOpe) int {
	if this == nil || this.db == nil {
		return -1
	}
	if nowDel.tableName == "" {
		return -2
	}

	commad := "select count(*) from " + nowDel.tableName

	rows, err := this.db.Query(commad)
	if checkError_Celica.CheckError(err, "Count record") == false {
		return -1
	}
	defer rows.Close()

	count := []byte{}
	for rows.Next() {
		if checkError_Celica.CheckError(rows.Scan(&count), "Count record") == false {
			return -1
		}
	}
	if checkError_Celica.CheckError(rows.Err(), "Count record") == false {
		return -1
	}

	return 0
}
