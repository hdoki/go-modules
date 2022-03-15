package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "somePassword"
	dbname   = "MyDB"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	// insert
	// hardcoded
	//insertStmt := `insert into "emp"("id", "name") values(7,'greg')`
	//_, e := db.Exec(insertStmt)
	//CheckError(e)

	// dynamic
	//insertDynStmt := `insert into "emp"("id", "name") values($1,$2)`
	//_, e := db.Exec(insertDynStmt, 8, "Jane")
	//CheckError(e)

	// update
	//updateStmt := `update "emp" set "id"=$1, "name"=$2 where "id"=$3`
	//_, e := db.Exec(updateStmt, 9, "Mary", 2)
	//CheckError(e)

	// Delete
	deleteStmt := `delete from "emp" where id=$1`
	_, e := db.Exec(deleteStmt, 9)
	CheckError(e)

	//Read data
	rows, err := db.Query(`SELECT "id", "name" FROM "emp" ORDER BY "id"`)
	countstmt := `select count(*) from emp`
	c := 0
	e = db.QueryRow(countstmt).Scan(&c)
	CheckError(e)
	fmt.Printf("Count is  %v \n", c)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string

		err = rows.Scan(&id, &name)
		CheckError(err)

		fmt.Println(id, name)
	}

	CheckError(err)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
