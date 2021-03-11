package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const username string = "root"
const password string = "11597"
const host string = "localhost"
const port int = 3306
const database string = "project_go_web"

var db *sql.DB

func CreateConnection() {
	if connection, err := sql.Open("mysql", generateURL()); err != nil {
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
		_, err := db.Exec(schema)
		if err != nil {
			log.Println(err)
		}
	}
}

func exitsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}
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

//<username>:<password>@tcp(<host>:<port>)/<database>
func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
}
