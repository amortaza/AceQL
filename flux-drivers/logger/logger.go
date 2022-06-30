package logger

import (
	"fmt"
	"time"
)

const (
	ERROR string = "ERROR"
)

func Log(msg string, source string) {
	//todo
	//fmt.Println(time.Now().Format(time.Kitchen), " (", source, ")", msg)
}

func Error(msg interface{}, source string) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println(time.Now().Format(time.Kitchen), " ***** ( ERROR ) ", source, msg)
}
