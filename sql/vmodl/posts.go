package vmodl

import (
	"github.com/google/uuid"
	"github.com/gowroc/meetups/sql/query"
	"github.com/gowroc/meetups/sql/user"
)

func (us *Users) InsertPost(up user.Post) error {
	return us.db.Insert(&up)
}

func (us *Users) GetUserWithPosts(id uuid.UUID) (user.UserWithPosts, error) {
	uwp := user.UserWithPosts{}
	err := us.db.SelectOne(&uwp, query.User.GetUserWithPosts, id)
	return uwp, err
}
