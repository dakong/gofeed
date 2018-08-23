package utils

import (
	"fmt"
	"os"
)

// HandleError ...
func HandleError(err error, msg string, exit bool) bool {
	if err != nil {
		fmt.Println(msg)
		if exit {
			os.Exit(1)
		}
		return true
	}

	return false
}
