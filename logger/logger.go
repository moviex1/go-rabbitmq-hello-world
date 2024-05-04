package logger

import "log"

func Info(msg string) {
	log.Printf("[info] : %s", msg)
}

func Error(msg string) {
	log.Fatalf("[error] : %s", msg)
}
