package logging

import (
	"fmt"
	"time"
)

func HandleErr(err error) {
	if err != nil {
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(nowTime, " [E] ", err)
	}
}

func Info(msg string) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(nowTime, " [I] ", msg)
}
