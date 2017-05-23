package user

import (
	"database/sql"

	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// User model.
type User struct {
	ID   uuid.UUID     `db:"user_id"`
	Name string        `db:"name"`
	Age  sql.NullInt64 `db:"age"`
}

type Post struct {
	ID      uuid.UUID `db:"user_post_id", json:"user_post:id"`
	UserID  uuid.UUID `db:"user_id", json:"user_id"`
	Content string    `db:"content", json:"content"`
}

type UserPosts []Post

func (up *UserPosts) Scan(i interface{}) error {
	bb, err := getBytes(i)
	if err != nil {
		return err
	}
	ps := []Post{}
	if err := json.Unmarshal(bb, &ps); err != nil {
		return err
	}
	*up = UserPosts(ps)
	return nil
}

type UserWithPosts struct {
	User
	UserPosts UserPosts `db:"posts"`
}

func getBytes(i interface{}) ([]byte, error) {
	switch t := i.(type) {
	case []byte:
		return t, nil
	case string:
		return []byte(t), nil
	default:
		return nil, fmt.Errorf("Can't convert %T to bytes", i)
	}
}
