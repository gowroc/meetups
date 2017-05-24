package vsqlx

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
	err := us.db.Get(&u, query.User.GetByID, id)
	return u, err
}

func (us *Users) GetAllUsers() ([]user.User, error) {
	users := []user.User{}
	err := us.db.Select(&users, query.User.SelectAll)
	return users, err

}

func (us *Users) DeleteAllUsers() error {
	_, err := us.db.Exec(query.User.DeleteAll)
	return err
}
