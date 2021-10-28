package logger

import "fmt"

type Source string

const (
	JsonEncoding Source = "JSON-ENCODING"
	MAIN = "MAIN"
)

func Log(msg string, source Source) {
	if source != JsonEncoding {
		fmt.Println("(", source, ")", msg)
	}
}

func Error(msg interface{}, source Source) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println("***** ( ERROR ) ", source, msg)
}

