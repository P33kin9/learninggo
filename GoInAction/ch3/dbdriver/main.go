package main

import (
	"database/sql"
	_ "goination/ch3/dbdriver/postgres"
)

func main() {
	sql.Open("postgres", "mydb")
}
