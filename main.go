package main 

import (
	"database/sql"
	"fmt"
  _"github.com/lib/pq"

)
func main () {

	connectionString :="host=localhost user=postgres password=postgres dbname=postgres sslmode=disable"
	
	db, err := sql.Open("postgres", connectionString)
	
	if err != nil {
        fmt.Println(err.Error())
    }
  
    defer db.Close()

    // Test the connection to the database
    if err := db.Ping(); err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully Connected")
    }

}
