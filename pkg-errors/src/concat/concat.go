package concat

import (
	"fmt"

	"github.com/gowroc/meetups/pkg-errors/src/common"
)

func CallA() error {
	return fmt.Errorf("Error from CallA: %v", CallB())
}

func CallB() error {
	return fmt.Errorf("Error from CallB: %v", CallC())
}

func CallC() error {
	return common.MyError{Msg: "Error from CallC"}
}
