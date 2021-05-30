package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

//ログの設定
func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func main() {
	//log
	LoggingSettings("test.log")
	_, err := os.Open("aaa")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	log.Fatalln("error!!")
	//Fatal以降は実行されない

	fmt.Println("ok!")
}