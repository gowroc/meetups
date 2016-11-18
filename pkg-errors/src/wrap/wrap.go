package wrap

import (
	"github.com/pkg/errors"

	"github.com/gowroc/meetups/pkg-errors/src/common"
)

func CallA() error {
	return errors.Wrap(CallB(), "Error from CallA")
}

func CallB() error {
	return errors.Wrap(CallC(), "Error from CallB")
}

func CallC() error {
	return common.MyError{Msg: "Error from CallC"}
}
