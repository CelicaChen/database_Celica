package database_Celica

import (
	"database/sql"
)

type CelicaSql struct {
	db *sql.DB
}

type RecordOpe struct {
	keyName   string
	keyValue  string
	tableName string
	Field     string
	NowData   string
}

func NewRecordOpe(tableName string, keyName string, keyValue string) *RecordOpe {
	tableName = "`" + tableName + "`"

	newOperating := new(RecordOpe)
	newOperating.tableName = tableName
	newOperating.keyName = keyName
	newOperating.keyValue = keyValue
	newOperating.Field = ""
	newOperating.NowData = ""

	return newOperating
}
