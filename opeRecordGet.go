package database_Celica

import (
	"Celica/checkError_Celica"
	"strconv"
)

func (this *CelicaSql) RecordGet(nowGet *RecordOpe) int {
	if this == nil || this.db == nil {
		return 1
	}
	if nowGet.tableName == "" || nowGet.Field == "" || nowGet.keyName == "" {
		return 2
	}

	commad := "select " + nowGet.Field + " from " + nowGet.tableName + " where " + nowGet.keyName + " =?"

	rows, err := this.db.Query(commad, nowGet.keyValue)
	if checkError_Celica.CheckError(err, "Get record") == false {
		return -1
	}
	defer rows.Close()

	in_data := ""
	for rows.Next() {
		if checkError_Celica.CheckError(rows.Scan(&in_data), "Get record") == false {
			return -1
		}
	}
	if checkError_Celica.CheckError(rows.Err(), "Get record") == false {
		return -1
	}

	nowGet.NowData = in_data
	return 0
}

func (this *CelicaSql) RecordGetInt(nowGet *RecordOpe) (int, int) {
	if this == nil || this.db == nil {
		return 1, 0
	}
	if nowGet.tableName == "" || nowGet.Field == "" || nowGet.keyName == "" {
		return 2, 0
	}

	commad := "select " + nowGet.Field + " from " + nowGet.tableName + " where " + nowGet.keyName + " =?"

	rows, err := this.db.Query(commad, nowGet.keyValue)
	if checkError_Celica.CheckError(err, "Get record") == false {
		return -1, 0
	}
	defer rows.Close()

	in_data := 0
	for rows.Next() {
		if checkError_Celica.CheckError(rows.Scan(&in_data), "Get record") == false {
			return -1, 0
		}
	}
	if checkError_Celica.CheckError(rows.Err(), "Get record") == false {
		return -1, 0
	}

	nowGet.NowData = strconv.Itoa(in_data)
	return 0, in_data
}
