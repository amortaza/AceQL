package logger

import (
	"fmt"
	"time"
)

type Source string

const (
	Bootstrap    Source = "Bootstrap"
	JsonEncoding        = "JSON-ENCODING"
	Main                = "Main"
	REST                = "REST"
	SQL                 = "SQL"
)

func Log(msg string, source Source) {
	fmt.Println(time.Now().Format(time.Kitchen), " (", source, ")", msg)
}

func Error(msg interface{}, source Source) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println(time.Now().Format(time.Kitchen), " ***** ( ERROR ) ", source, msg)
}
