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

func AstNightsToMin(AstTime time.Time, nights int) float64 {
	// アストルティアの66夜後が何分後になるか
	targetTime := AstTime.AddDate(0, 0, nights)
	duration := targetTime.Sub(AstTime)
	return duration.Minutes() / 20
}

func main() {
	daySec := DaySeconds(time.Now())
	fmt.Println(daySec)
	Adate := SecToAstoltiaTime(180)
	fmt.Println("Astoltia time is ", Adate)
	tomin := AstNightsToMin(Adate, 66)
	fmt.Println(tomin)
	fmt.Println(time.Now().Add(time.Duration(tomin) * time.Minute))

	fmt.Println("Astoltia time is ", Adate)
	fmt.Println("Astoltia time is ", SecToAstoltiaTime(int(daySec)))

	ColossusTime := time.Date(2020, 3, 10, 7, 0, 0, 0, time.Local)
	day, hour, min, sec := AstoltiaNights(ColossusTime)
	fmt.Printf("Colossus time is %dNights %2d:%2d:%2d\n", day, hour, min, sec)
	fmt.Println(ColossusTime.Add(time.Duration(AstNightsToMin(ColossusTime, 1)) * time.Minute))

	day, hour, min, sec = AstoltiaNights(time.Now())
	fmt.Printf("Astoltia time is %dNights %2d:%2d:%2d\n", day, hour, min, sec)
}
