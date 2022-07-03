package logger

import (
	"errors"
	"fmt"
	"time"
)

type Source string

const (
	Bootstrap    Source = "Bootstrap"
	ERROR               = "ERROR"
	JsonEncoding        = "JSON-ENCODING"
	Main                = "Main"
	REST                = "REST"
	SQL                 = "SQL"
)

func Log(msg string, source Source) {
	fmt.Println(time.Now().Format(time.Kitchen), " (", source, ")", msg)
}
func Err(err error, source Source) error {
	return Error(err.Error(), source)
}

func Error(msg string, source Source) error {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	fmt.Println(time.Now().Format(time.Kitchen), " ***** ( ERROR ) ", source, msg)

	return errors.New(msg)
}
