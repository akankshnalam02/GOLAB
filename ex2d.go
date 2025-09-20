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
		log.Fatal("Failed to connect:", err)
	}
	fmt.Println("Connected to PostgreSQL")
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT UNIQUE)`)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = db.Exec(`INSERT INTO users (name, email) VALUES ($1, $2)`, "John Doe", "john@example.com")
	rows, _ := db.Query("SELECT id, name, email FROM users")
	defer rows.Close()
	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		fmt.Println(id, name, email)
	}
	_, _ = db.Exec(`UPDATE users SET email=$1 WHERE name=$2`, "john.doe@example.com", "John Doe")
	_, _ = db.Exec(`DELETE FROM users WHERE name=$1`, "John Doe")
}
