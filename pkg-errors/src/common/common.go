package common

import "fmt"

type MyError struct {
	Msg string
}

func (me MyError) Error() string {
	return fmt.Sprintf("Errrrr: %s", me.Msg)
}
