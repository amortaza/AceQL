package logger

import (
	"fmt"
	"os"
	"time"
)

type Logfile struct {
	filename string
	file     *os.File
}

func NewLogfile(filename string) (*Logfile, error) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("***** ( ERROR ) could not create log file, " + err.Error())
		return nil, err
	}

	//info, err := file.Stat()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return nil, err
	//}
	//
	//fmt.Println(info.Mode().IsDir())

	return &Logfile{filename: filename, file: file}, nil
}

func (logfile *Logfile) Write(entrytype, source, msg string) (string, error) {
	timestr := time.Now().Format("2006-01-02 03:04:05")

	entry := fmt.Sprintf("%s\t%s\t%s\t%s\n", timestr, entrytype, source, msg)

	var err error
	if _, err = logfile.file.WriteString(entry); err != nil {
		fmt.Println("***** ( ERROR ) failed to write entry to log file, " + err.Error())
	}

	return entry, err
}
