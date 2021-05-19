package console_util

import "fmt"

func Input(msg string) string {
	value := ""
	fmt.Print(msg + ": ")
	fmt.Scan(&value)
	return value
}
