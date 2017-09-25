package actions

import (
	"errors"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"

	"github.com/gowroc/meetups/buffalo/gowroc/models"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Error(500, errors.New(""))
	}

	var ps []models.Post
	if err := tx.All(&ps); err != nil {
		return c.Error(500, errors.New(""))
	}
	c.Set("posts", ps)

	return c.Render(200, r.HTML("index.html"))
}
