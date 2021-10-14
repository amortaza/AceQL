package logger

import "fmt"

func Log(msg string, source string) {
	fmt.Println("(", source, ")", msg)
}

func Error(msg interface{}, source string) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println("***** ( ERROR ) ", source, msg)
}

