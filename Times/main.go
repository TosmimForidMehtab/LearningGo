package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)                                      // 2024-05-03 16:05:11.9406107 +0530 IST m=+0.003385701
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday")) // 05-03-2024 16:07:58 Friday

	createdDate := time.Date(2024, time.January, 10, 12, 10, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
