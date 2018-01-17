package util

import "fmt"

func LogErr(err error) bool{
	if err != nil {
		fmt.Println(err)
	}
	return err != nil
}