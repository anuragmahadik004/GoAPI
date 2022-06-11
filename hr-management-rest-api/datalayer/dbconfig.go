package datalayer

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server   = "CRYPTON\\MSSQLSERVER2019"
	user     = "sa"
	password = "pass@123"
	database = "HR_Database"
)

var connString = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", server, user, password, database)

func GetDbCOnnection() (conn *sql.DB, err error) {

	conn, err = sql.Open("mssql", connString)

	return
}
