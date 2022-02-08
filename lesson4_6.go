package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func main() {
	LoggingSettings("test.log")

	_, err := os.Open("hogehoge")
	if err != nil {
		log.Fatalln("v ...interface{}", err)
	}

	log.Println("logging")
	log.Printf("%T, %v", "test", "log")

	log.Fatalf("%T, %v", "error", "golang")
	log.Fatalln("panick!")

	fmt.Println("hogehoge")
}
