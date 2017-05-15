package database_Celica

import (
	"Celica/checkError_Celica"
)

func (this *CelicaSql) RecordSet(nowSet *RecordOpe) int {
	if this == nil || this.db == nil {
		return 1
	}
	if nowSet.tableName == "" || nowSet.Field == "" || nowSet.keyName == "" {
		return 2
	}

	commad := "update " + nowSet.tableName + " set " + nowSet.Field + " =? where " + nowSet.keyName + " =?"

	stmt, err := this.db.Prepare(commad)
	if checkError_Celica.CheckError(err, "Set record") == false {
		return -1
	}
	defer stmt.Close()

	result, err := stmt.Exec(nowSet.NowData, nowSet.keyValue)
	if checkError_Celica.CheckError(err, "Set record") == false {
		return -1
	}

	_, err = result.RowsAffected()
	if checkError_Celica.CheckError(err, "Set record") == false {
		return -1
	}

	return 0
}

func (this *CelicaSql) RecordSetInt(nowSet *RecordOpe, nowData int) int {
	if this == nil || this.db == nil {
		return 1
	}
	if nowSet.tableName == "" || nowSet.Field == "" || nowSet.keyName == "" {
		return 2
	}

	commad := "update " + nowSet.tableName + " set " + nowSet.Field + " =? where " + nowSet.keyName + " =?"

	stmt, err := this.db.Prepare(commad)
	if checkError_Celica.CheckError(err, "Set record") == false {
		return -1
	}
	defer stmt.Close()

	result, err := stmt.Exec(nowData, nowSet.keyValue)
	if checkError_Celica.CheckError(err, "Set record") == false {
		return -1
	}

	_, err = result.RowsAffected()
	if checkError_Celica.CheckError(err, "Set record") == false {
		return -1
	}

	return 0
}

func (this *CelicaSql) RecordIntChange(nowSet *RecordOpe, change int) int {
	reply, data := this.RecordGetInt(nowSet)
	if reply != 0 {
		return reply
	}

	data += change
	reply = this.RecordSetInt(nowSet, data)
	if reply != 0 {
		return reply
	}

	return 0
}
