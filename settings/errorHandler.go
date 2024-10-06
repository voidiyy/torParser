package settings

import (
	"fmt"
	"os"
)

func ErrorExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
