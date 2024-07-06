package logger

import "fmt"

func Log(msg string) {
	fmt.Printf("[INFO] %s\n", msg)
}

func Error(msg string) {
	fmt.Printf("[ERROR] %s\n", msg)
}
