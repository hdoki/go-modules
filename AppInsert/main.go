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
	//connection string
	dbconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//open database
	db, err := sql.Open("postgres", dbconn)
	CheckError(err)

	//close databse
	defer db.Close()

	//check database
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected")

	n := 0
	fmt.Println("Enter choice: 1 for Select, 2 for Insert, 3 for Update, 4 for Delete")
	fmt.Scan(&n)
	selection := true
	for selection {
		selection = false
		switch n {
		case 1: //select
			rows, err := db.Query(`SELECT "id", "name" FROM "emp" ORDER BY "id"`)
			CheckError(err)
			defer rows.Close()
			for rows.Next() {
				var id int
				var name string

				err = rows.Scan(&id, &name)
				CheckError(err)
				fmt.Println(id, name)
			}
			countstmt := `select count(*) from emp`
			c := 0
			e := db.QueryRow(countstmt).Scan(&c)
			CheckError(e)
			fmt.Printf("Count is  %v \n", c)
			break

		case 2: //insert
			sqlstmt := `insert into "emp"("id","name") values(1,'aaa')`
			_, err := db.Exec(sqlstmt)
			CheckError(err)
			fmt.Println("Successfully Inserted Row")
		case 3: //Update
			sqlstmt := `update "emp" set "id"=$1, "name"=$2 where "id"=$3`
			_, e := db.Exec(sqlstmt, 9, "Mary", 1)
			CheckError(e)
			fmt.Println("Update Successful")
		case 4: //Delete
			sqlstmt := `delete from "emp" where "id"=$1`
			_, e := db.Exec(sqlstmt, 9)
			CheckError(e)
			fmt.Println("Row deleted")

		default:
			fmt.Println("Invalid choice")
			selection = false
		}
	}
}
func CheckError(err error) {

	if err != nil {
		panic(err)
	}
}
