package logger

import (
	"errors"
	"fmt"
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

var g_logfile *Logfile

func init() {
	fmt.Println("initializing bsnlogs.log")

	var err error
	g_logfile, err = NewLogfile("bsnlogs.log")
	if err != nil {
		fmt.Println("***** ( ERROR ) could not initialize bsnlogs.log, see " + err.Error())
	} else {
		fmt.Println("successfully initialized bsnlogs.log")
	}
}

func Log(msg string, source Source) {
	// fmt.Println(/*time.Now().Format(time.Kitchen),*/ "(", source, ")", msg)
	g_logfile.Write("info", string(source), msg)
}
func Err(err error, source Source) error {
	Error(err.Error(), source)
	return err
}

func Error(msg string, source Source) error {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println( /*time.Now().Format(time.Kitchen), */ "***** ( ERROR )", source, msg)

	g_logfile.Write("error", string(source), msg)

	return errors.New(msg)
}
