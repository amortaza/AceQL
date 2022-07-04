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
	GRPC                = "GRPC"
	JsonEncoding        = "JSON-ENCODING"
	REST                = "REST"
	SQL                 = "SQL"
)

func Log(msg string, source Source) {
	fmt.Println(time.Now().Format(time.Kitchen), " (", source, ")", msg)
}
func Err(err error, source Source) error {
	Error(err.Error(), source)
	return err
}

func Error(msg string, source Source) error {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	fmt.Println(time.Now().Format(time.Kitchen), " ***** ( ERROR ) ", source, msg)

	return errors.New(msg)
}
