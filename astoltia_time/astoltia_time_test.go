package main

import (
	"testing"
	"time"
)

func TestDaySeconds(t *testing.T) {
	result := DaySeconds(time.Now())
	if int(result) <= 0 {
		t.Errorf("DaySeconds(time.Now()) = %v, expected 0", result)
	}
}
func TestSecToAstoltiaTime(t *testing.T) {
	//var result int

	Adate := SecToAstoltiaTime(180)
	if Adate.Hour() != 1 {
		t.Errorf("time to failed. exect:%d, actual:%d", 1, Adate.Hour())
	}

	//loging
	t.Logf("Adate.Hour is %d", Adate.Hour())

}
