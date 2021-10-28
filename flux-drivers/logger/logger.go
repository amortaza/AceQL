package logger

import "fmt"

const (
	ERROR string = "ERROR"
)

func Log(msg string, source string) {
	fmt.Println("(", source, ")", msg)
}

func Error(msg interface{}, source string) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println("***** ( ERROR ) ", source, msg)
}

