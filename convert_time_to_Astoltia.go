package main

import (
	"fmt"
	"time"
)

func ConvertToAstoltiaTime(setTime time.Time) float64 {
	y := setTime.Year()
	m := setTime.Month()
	d := setTime.Day()
	initDay := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	sub := setTime.Sub(initDay)
	fmt.Println(y, m, d, initDay, sub)
	AstoltiaTime := sub.Minutes() / 72
	return AstoltiaTime
}

func main() {
	Atime := ConvertToAstoltiaTime(time.Now())
	fmt.Printf("Astoltia time is %f.\n", Atime)

}
