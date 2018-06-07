package Service

import (
	"fmt"
)

func checkerr(err error, name string) {
	if err != nil {
		fmt.Printf("err at %s \n The err is %s \n", name, err)
	}
}
