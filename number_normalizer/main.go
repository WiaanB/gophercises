package main

import (
	"fmt"
	"regexp"

	database "workspace/number_normalizer/db"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "client_pass"
	dbname   = "gophercises"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

	must(database.Reset("postgres", fmt.Sprintf("%s dbname=%s", psqlInfo, "postgres"), dbname))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err := database.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(database.Migrate("postgres", fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)))

	err = db.Seed()
	must(err)

	phones, err := db.AllPhones()
	must(err)

	for _, p := range phones {
		fmt.Printf("working on... %v\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			fmt.Println("updating or removing...", number)
			existing, err := db.FindNumber(number)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p.ID))
			} else {
				p.Number = number
				must(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("no changes required")
		}
	}
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
