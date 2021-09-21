package logger

import "fmt"

type Source string

const (
	SQL Source = "SQL"
	Bootstrap = "Bootstrap"
	Main = "Main"
	REST = "REST"
)

func Log(msg string, source Source) {
	fmt.Println("(", source, ")", msg)
}

func Error(msg interface{}, source Source) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println("***** ( ERROR ) ", source, msg)
}

