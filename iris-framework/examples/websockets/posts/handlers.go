package posts

import (
	"database/sql"
	"fmt"
)

func GetAll(db *sql.DB) ([]Person, error) {
	rows, err := db.Query("SELECT * FROM people")
	if err != nil {
		return nil, fmt.Errorf("failed to get people list: %v", err)
	}

	var people []Person
	for rows.Next() {
		p, err := readPerson(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to get people list: %v", err)
		}

		people = append(people, p)
	}

	return people, nil
}

func Get(db *sql.DB, ID int) (Person, error) {
	rows, err := db.Query("SELECT * FROM people WHERE id = $1", ID)
	rows.Next()
	if err != nil {
		return Person{}, fmt.Errorf("failed to get people list: %v", err)
	}

	return readPerson(rows)
}

func Post(db *sql.DB, name, hobby string) error {
	_, err := db.Query("INSERT INTO people (name, hobby) VALUES ($1, $2)", name, hobby)
	return err
}

func Delete(db *sql.DB, ID int) error {
	_, err := db.Query("DELETE * FROM people WHERE id = $1", ID)
	return err
}

func readPerson(rows *sql.Rows) (Person, error) {
	var id int
	var name string
	var hobby string
	err := rows.Scan(&id, &name, &hobby)
	if err != nil {
		return Person{}, fmt.Errorf("failed to get people list: %v", err)
	}

	return Person{id, name, hobby}, nil
}

type Person struct {
	ID    int
	Name  string
	Hobby string
}
