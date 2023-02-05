package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql, Open("mysql", "root:gRY0s%a5RT0l@/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	// _, err = stmt.Exec(1, "Tiago") // chave duplicada

	if err != nil {
		tx.Rollblack()
		log.Fatal(err)
	}

	tx.Commit()

}
