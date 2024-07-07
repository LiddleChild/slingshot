package logger

import "fmt"

func Log(msg string) {
	fmt.Println(msg)
}

func Error(msg string) {
	fmt.Printf("error: %s\n", msg)
}
