package logger

import (
	"errors"
	"fmt"
	"time"
)

func PushStackTrace(s string, err error) error {
	return errors.New(s + " > " + err.Error())
}

func Info(msg string, source string) {
	entry := formatLog("info", string(source), msg)

	fmt.Println(entry)
}

func Err(err error, source string) error {
	Error(err.Error(), source)
	return err
}

func Error(msg string, source string) error {
	entry := formatLog("error", string(source), msg)

	fmt.Println(entry)

	return errors.New(msg)
}

func formatLog(entrytype, source, msg string) string {
	timestr := time.Now().Format("2006-01-02 15:04:05")

	entry := fmt.Sprintf("%s\t%s\t%s\t%s\n", timestr, entrytype, source, msg)

	return entry
}
