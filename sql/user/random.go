package user

import (
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

// RandomUser which can be inserted to DB.
func RandomUser() User {
	u := User{
		ID:   uuid.New(),
		Name: fmt.Sprintf("name-%d", rand.Intn(10000)),
	}
	age := rand.Intn(100) - 50
	if age > 0 {
		u.Age = sql.NullInt64{
			Valid: true,
			Int64: int64(age),
		}
	}
	return u
}

// RandomPost which can be inserted to DB.
func RandomPost(uid uuid.UUID) Post {
	return Post{
		ID:      uuid.New(),
		UserID:  uid,
		Content: fmt.Sprintf("post-%d", rand.Intn(10000)),
	}
}
