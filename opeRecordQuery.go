package database_Celica

import (
	"Celica/checkError_Celica"
)

func (this *CelicaSql) RecordQuery(newQuery *RecordOpe) []string {
	list := []string{}
	if this == nil || this.db == nil {
		return list
	}
	if newQuery.tableName == "" || newQuery.keyName == "" {
		return list
	}

	commad := "select " + newQuery.keyName + " from " + newQuery.tableName

	row, err := this.db.Query(commad)
	if checkError_Celica.CheckError(err, "Query compID") == false {
		return list
	}
	defer row.Close()

	temp := ""
	for i := 0; row.Next(); i++ {
		row.Scan(&temp)
		list = append(list, temp)
	}

	return list
}

func (this *CelicaSql) RecordQuerySorting(newQuery *RecordOpe, filter []string) []string {
	list := []string{}
	if this == nil || this.db == nil {
		return list
	}
	if newQuery.tableName == "" || newQuery.keyName == "" {
		return list
	}

	commad := "select " + newQuery.keyName + " from " + newQuery.tableName
	length := len(filter)
	if length != 0 {
		commad += " where " + filter[0]
		for i := 1; i < length; i++ {
			commad += " and " + filter[i]
		}
	}
	if newQuery.Field != "" {
		commad += " order by " + newQuery.Field
		if newQuery.NowData == "desc" {
			commad += " desc"
		} else {
			commad += " asc"
		}
	}

	row, err := this.db.Query(commad)
	if checkError_Celica.CheckError(err, "Query compID") == false {
		return list
	}
	defer row.Close()

	temp := ""
	for i := 0; row.Next(); i++ {
		row.Scan(&temp)
		list = append(list, temp)
	}

	return list
}
