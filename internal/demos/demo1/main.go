package main

import (
	"fmt"
	"time"

	"github.com/go-zwbc/timezh"
)

func init() {
	time.Local = timezh.TZ.CST //通过这样，也能改变整个程序的运行时区位置，但不建议这样做
}

func main() {
	now := time.Now()
	fmt.Println(now)

	tmz := now.Location()
	fmt.Println(tmz.String())
}
