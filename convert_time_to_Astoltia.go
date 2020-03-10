package main

import (
	"fmt"
	"time"
)

func DaySeconds(setTime time.Time) float64 {
	duration := setTime.Sub(time.Date(setTime.Year(), setTime.Month(), setTime.Day(), 0, 0, 0, 0, time.Local))
	return duration.Seconds()
}

func SecToAstoltiaTime(sec int) time.Time {
	astSec := sec * 20
	astDate := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)
	return astDate.Add(time.Duration(astSec) * time.Second)
}

func AstoltiaNights(setTime time.Time) (astDay int, astHour int, astMin int, astSec int) {
	daySec := DaySeconds(setTime)
	AstTime := SecToAstoltiaTime(int(daySec))
	return AstTime.Day(), AstTime.Hour(), AstTime.Minute(), AstTime.Second()
}

func main() {
	daySec := DaySeconds(time.Now())
	fmt.Println(daySec)
	Adate := SecToAstoltiaTime(180)
	fmt.Println("Astoltia time is ", Adate)
	fmt.Println("Astoltia time is ", SecToAstoltiaTime(int(daySec)))
	day, hour, min, sec := AstoltiaNights(time.Now())
	fmt.Printf("Astoltia time is %dNights %d:%d:%d", day, hour, min, sec)
}
