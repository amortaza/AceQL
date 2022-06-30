package logger

import (
	"fmt"
	"time"
)

type Source string

const (
	SQL       Source = "SQL"
	Bootstrap        = "Bootstrap"
	Main             = "Main"
	REST             = "REST"
)

func Log(msg string, source Source) {
	//todo
	// fmt.Println(time.Now().Format(time.Kitchen), " (", source, ")", msg)
}

func Error(msg interface{}, source Source) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println(time.Now().Format(time.Kitchen), " ***** ( ERROR ) ", source, msg)
}
