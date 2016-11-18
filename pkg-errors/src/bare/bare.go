package bare

import "github.com/gowroc/meetups/pkg-errors/src/common"

func CallA() error {
	return CallB()
}

func CallB() error {
	return CallC()
}

func CallC() error {
	return common.MyError{Msg: "Error from CallC"}
}
