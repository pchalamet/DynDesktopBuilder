package core

import "fmt"

func CheckError(err error, description string) {
	if err != nil {
		msg := fmt.Sprintf("%s \"%s\"", description, err)
		panic(msg)
	}
}
