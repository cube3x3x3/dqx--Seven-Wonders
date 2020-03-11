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

func AstNightsToMin(setTime time.Time, nights int) float64 {
	// アストルティアのn夜後が、リアル時間のn分後になるか
	// アストルティアは20倍速で時間が進んでいる
	targetTime := setTime.AddDate(0, 0, nights)
	duration := targetTime.Sub(setTime)
	return duration.Minutes() / 20
}

func getNextWonderTime(setTime time.Time, nights int) time.Time {
	tomin := AstNightsToMin(setTime, nights)
	return setTime.Add(time.Duration(tomin) * time.Minute)
}

func main() {
	day, hour, min, sec := AstoltiaNights(time.Now())
	fmt.Printf("Astoltia time is %dNights %2d:%2d:%2d\n", day, hour, min, sec)
}
