package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "client_pass"
	dbname   = "gophercises"
)

var numbers []string = []string{"1234567890", "123 456 7891", "(123) 456 7892", "(123) 456-7893", "123-456-7894", "123-456-7890", "1234567892", "(123)456-7892"}

type phone struct {
	id     int
	number string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

	db, err := sql.Open("postgres", fmt.Sprintf("%s dbname=%s", psqlInfo, "postgres"))
	must(err)
	err = resetDB(db, dbname)
	must(err)
	db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	// add all numbers
	must(createNumberTable(db))
	for _, num := range numbers {
		_, err = insertPhone(db, num)
		must(err)
	}

	// select single number
	number, err := getNumber(db, 2)
	must(err)
	fmt.Printf("Second record is: %v\n", number)

	// select all numbers
	numbers, err := allNumbers(db)
	must(err)
	for _, p := range numbers {
		fmt.Printf("working on... %v\n", p)
		number := normalize(p.number)
		if number != p.number {
			fmt.Println("updating or removing...", number)
			existing, err := findNumber(db, number)
			must(err)
			if existing != nil {
				must(deletePhone(db, p.id))
			} else {
				p.number = number
				must(updatePhone(db, p))
			}
		} else {
			fmt.Println("no changes required")
		}
	}
}

func deletePhone(db *sql.DB, id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.Exec(statement, id)
	return err
}

func updatePhone(db *sql.DB, p phone) error {
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(statement, p.id, p.number)
	return err
}

func findNumber(db *sql.DB, number string) (*phone, error) {
	var p phone
	err := db.QueryRow("SELECT id, value FROM phone_numbers WHERE value=$1", number).Scan(&p.id, &p.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func getNumber(db *sql.DB, id int) (phone, error) {
	var p phone
	err := db.QueryRow("SELECT id, value FROM phone_numbers WHERE id=$1", id).Scan(&p.id, &p.number)
	if err != nil {
		return p, err
	}
	return p, nil
}

func allNumbers(db *sql.DB) ([]phone, error) {
	var phones []phone
	rows, err := db.Query("SELECT id, value FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p phone
		if err := rows.Scan(&p.id, &p.number); err != nil {
			return nil, err
		}
		phones = append(phones, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return phones, nil
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func createNumberTable(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_numbers (
			id SERIAL,
			value VARCHAR(255)
		)`
	_, err := db.Exec(statement)
	return err
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	must(err)
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	must(err)
	return nil
}

func normalize(phone string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(phone, "")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
