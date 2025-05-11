package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const maxOpenConns = 5
const maxIdleConns = 5

var connection sql.DB
var inited = false

func GetConnection() *sql.DB {
	pass := "root"
	login := "root"
	port := "3306"
	host := "127.0.0.1"
	name := "4devbau"

	var connectStr = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", login, pass, host, port, name)
	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		log.Fatal("database connection error", err)
	}

	inited = true
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func ConvertSqlStringToString(valueSql sql.NullString) string {
	value := ""
	if valueSql.Valid {
		value = valueSql.String
	}

	return value
}

func ConvertSqlInt32ToInt(valueSql sql.NullInt32) int {
	value := 0
	if valueSql.Valid {
		value = int(valueSql.Int32)
	}

	return value
}

func ConvertToInterfaceSlice(ints []int) []interface{} {
	result := make([]interface{}, len(ints))
	for i, v := range ints {
		result[i] = v
	}
	return result
}
