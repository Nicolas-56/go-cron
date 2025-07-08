package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"home.ru/internal/config"
)

const maxOpenConns = 5
const maxIdleConns = 5

var connectionDB *DB
var inited = false

type DB struct {
	*sql.DB
}

func GetConnection() (*DB, error) {

	if !inited {
		configDB := config.GetConfig()
		var connectStr = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", configDB.Database.DbUser, configDB.Database.DbPass, configDB.Database.DbHost, configDB.Database.DbPort, configDB.Database.DbName)
		db, err := sql.Open("mysql", connectStr)
		if err != nil {
			log.Fatal("database connection error", err)
		}
		connectionDB = &DB{db}

		inited = true
		db.SetMaxOpenConns(maxOpenConns)
		db.SetMaxIdleConns(maxIdleConns)
		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}
	}

	return connectionDB, nil
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
