package database_Celica

import (
	"Celica/checkError_Celica"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Open(sqlAdress string) (*CelicaSql, error) {
	db, err := sql.Open("mysql", sqlAdress)
	if checkError_Celica.CheckError(err, "Open database") == false {
		return nil, err
	}
	newSql := new(CelicaSql)
	newSql.db = db

	return newSql, nil
}

func Close(thisSql *CelicaSql) int {
	if thisSql == nil || thisSql.db == nil {
		return 1
	}
	thisSql.db.Close()

	return 0
}
