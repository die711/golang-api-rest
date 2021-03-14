package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"rest/config"
)

var db *sql.DB
var debug bool

func init() {
	CreateConnection()
	debug = config.GetDebug()
	CreateTables()
}

func GetConnection() *sql.DB {
	return db
}

func InsertData(query string, arg ...interface{}) (int64, error) {
	if result, err := Exec(query, arg...); err != nil {
		return 0, err
	} else {
		id, err := result.LastInsertId()
		return id, err
	}
}

func CreateConnection() {

	if GetConnection() != nil {
		return
	}

	url := config.GetUrlDatabase()
	if connection, err := sql.Open("mysql", url); err != nil {
		panic(err)
	} else {
		db = connection
		fmt.Println("Conexion exitosa!")
	}
}
func CreateTables() {
	createTable("users", UserSchema)
}

func createTable(tableName, schema string) {
	if !exitsTable(tableName) {
		Exec(schema)
	} else {
		truncateTable(tableName)
	}
}

func truncateTable(tableName string) {
	sql := fmt.Sprintf("truncate %s", tableName)
	Exec(sql)
}

func exitsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, _ := Query(sql)
	return rows.Next()
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func CloseConnection() {
	db.Close()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil && !debug {
		log.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)

	if err != nil && !debug {
		log.Println(err)
	}
	return rows, err

}
