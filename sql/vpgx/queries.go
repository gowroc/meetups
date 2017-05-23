package vpgx

import (
	"github.com/gowroc/meetups/sql/query"
	"github.com/gowroc/meetups/sql/user"

	// required to set up the driver
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (us *Users) InsertUser(u user.User) error {
	_, err := us.db.Exec(query.User.Insert, u.ID, u.Name, u.Age)
	return err
}

func (us *Users) GetUser(id uuid.UUID) (user.User, error) {
	row := us.db.QueryRow("getByID", id)
	u := user.User{}
	err := row.Scan(&u.ID, &u.Name, &u.Age)
	return u, err
}

func (us *Users) GetAllUsers() ([]user.User, error) {
	rows, err := us.db.Query("selectAll")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := []user.User{}
	for rows.Next() {
		u := user.User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

func (us *Users) DeleteAllUsers() error {
	_, err := us.db.Exec(query.User.DeleteAll)
	return err
}
