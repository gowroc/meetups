package vsql

import (
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/gowroc/meetups/sql/query"
	"github.com/gowroc/meetups/sql/user"
)

func (us *Users) InsertUser(u user.User) error {
	_, err := us.db.Exec(query.User.Insert, u.ID, u.Name, u.Age)
	return err
}

func (us *Users) GetUser(id uuid.UUID) (user.User, error) {
	u := user.User{}
	err := us.db.QueryRow(query.User.GetByID, id).Scan(&u.ID, &u.Name, &u.Age)
	return u, err
}

func (us *Users) GetAllUsers() ([]user.User, error) {
	users := []user.User{}
	rows, err := us.db.Query(query.User.SelectAll)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		u := user.User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return users, err
		}
		users = append(users, u)
	}
	return users, nil

}

func (us *Users) DeleteAllUsers() error {
	_, err := us.db.Exec(query.User.DeleteAll)
	return err
}
