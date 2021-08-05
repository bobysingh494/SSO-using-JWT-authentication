package config

import(
	//"net/http"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DBconnection() *sql.DB{
	db, err := sql.Open("mysql", "root:Boby@1611@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	return db
}