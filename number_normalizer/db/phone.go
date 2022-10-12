package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DataBase struct {
	DB *sql.DB
}

type Phone struct {
	ID     int
	Number string
}

func (db *DataBase) Close() error {
	return db.DB.Close()
}

func Open(driverName, dataSource string) (*DataBase, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &DataBase{db}, nil
}

func Reset(driverName, dataSource, dbName string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	resetDB(db, dbName)
	return db.Close()
}

func Migrate(driverName, dataSource string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = createNumberTable(db)
	if err != nil {
		return err
	}
	return db.Close()
}

func (db *DataBase) Seed() error {
	var numbers []string = []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	// add all numbers
	for _, num := range numbers {
		_, err := insertPhone(db.DB, num)
		if err != nil {
			return err
		}
	}
	return nil
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
	if err != nil {
		return err
	}
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
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

func (db *DataBase) AllPhones() ([]Phone, error) {
	var phones []Phone
	rows, err := db.DB.Query("SELECT id, value FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Phone
		if err := rows.Scan(&p.ID, &p.Number); err != nil {
			return nil, err
		}
		phones = append(phones, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return phones, nil
}

func (db *DataBase) FindNumber(number string) (*Phone, error) {
	var p Phone
	err := db.DB.QueryRow("SELECT id, value FROM phone_numbers WHERE value=$1", number).Scan(&p.ID, &p.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (db *DataBase) DeletePhone(id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.DB.Exec(statement, id)
	return err
}

func (db *DataBase) UpdatePhone(p *Phone) error {
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.DB.Exec(statement, p.ID, p.Number)
	return err
}
