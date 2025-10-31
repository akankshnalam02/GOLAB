package main

import (
"database/sql"
"fmt"
"log"
_ "github.com/lib/pq"
)

func main() {
connStr := "user=username password='abcd123' dbname=my_database sslmode=disable"
db, err := sql.Open("postgres", connStr)
if err != nil {
log.Fatal(err)
}
defer db.Close()

err = db.Ping()
if err != nil {
log.Fatal("Failed to connect to the database:", err)
}

fmt.Println("Connected to PostgreSQL!")
}
