package logging

import (
	"log"
	"time"
)

func HandleErr(err error) {
	if err != nil {
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		log.Println(nowTime, " [E] ", err)
	}
}

func Info(msg string) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	log.Println(nowTime, " [I] ", msg)
}