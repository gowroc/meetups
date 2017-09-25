package grifts

import (
	"fmt"

	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("gowroc", func() {

	grift.Desc("hello", "greets the GoWroc community")
	grift.Add("hello", func(c *grift.Context) error {
		fmt.Println("Hello GoWroc!")
		return nil
	})

})
