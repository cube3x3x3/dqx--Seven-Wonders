package main

import (
	"fmt"
	"time"
)

func DayMinutes(setTime time.Time) float64 {
	duration := setTime.Sub(time.Date(setTime.Year(), setTime.Month(), setTime.Day(), 0, 0, 0, 0, time.Local))
	return duration.Minutes()
}

func ConvertToAstoltiaTime(setTime time.Time) float64 {
	sub := DayMinutes(setTime)
	AstoltiaDays := int(sub / 72)
	AstoltiaTime := sub - (float64(AstoltiaDays) * 72)
	AstHour := int(AstoltiaTime)
	AstMin := (AstoltiaTime - float64(AstHour)) * 60
	fmt.Printf("hour %d, min %f", AstHour, AstMin)
	return AstoltiaTime
}

func main() {
	Atime := ConvertToAstoltiaTime(time.Now())
	fmt.Printf("Astoltia time is %f.\n", Atime)

}
