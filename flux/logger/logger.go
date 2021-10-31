package logger

import (
	"fmt"
	"time"
)

type Source string

const (
	JsonEncoding Source = "JSON-ENCODING"
	MAIN = "MAIN"
)

func Log(msg string, source Source) {
	if source == JsonEncoding {
		return
	}

	fmt.Println(time.Now().Format(time.Kitchen), " (", source, ")", msg)
}

func Error(msg interface{}, source Source) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println(time.Now().Format(time.Kitchen), " ***** ( ERROR ) ", source, msg)
}

