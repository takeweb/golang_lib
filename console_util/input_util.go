package console_util

import "fmt"

func Input(msg string) string {
	var value string
	fmt.Print(msg + ": ")
	fmt.Scan(&value)
	return value
}
