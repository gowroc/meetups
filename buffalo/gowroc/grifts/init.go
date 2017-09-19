package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gowroc/meetups/buffalo/gowroc/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
