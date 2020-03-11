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
func TestAstNightsToMin(t *testing.T) {
	Adate := SecToAstoltiaTime(0)
	tomin := AstNightsToMin(Adate, 66)
	if tomin != 4752 {
		t.Errorf("AstNightsToMin to faild. exect:%d, actual:%f", 4752, tomin)
	}
	//loging
	t.Logf("toMin is %f", tomin)

}

func TestAstoltiaNights(t *testing.T) {
	colossusTime := time.Date(2020, 3, 10, 7, 0, 0, 0, time.Local)
	day, hour, min, sec := AstoltiaNights(colossusTime)
	if day != 6 || hour != 20 || min != 0 || sec != 0 {
		t.Errorf("TestAstoltiaNights to faild. exect:%d %d %d %d, actual:%d %d %d %d", 6, 20, 0, 0, day, hour, min, sec)
	}
}

func TestAstNightsToReal(t *testing.T) {
	colossusTime := time.Date(2020, 3, 10, 7, 0, 0, 0, time.Local)
	tomin := AstNightsToMin(colossusTime, 66)
	result := colossusTime.Add(time.Duration(tomin) * time.Minute)
	if result.Year() != 2020 || result.Month() != 3 || result.Day() != 13 || result.Hour() != 14 || result.Minute() != 12 || result.Second() != 0 {
		t.Errorf("TestAstoltiaNights to faild. exect:2020-3-13 14:12:0, actual:%d-%d-%d %d:%d:%d", result.Year(), result.Month(), result.Day(), result.Hour(), result.Minute(), result.Second())
	}
}

func TestAstNightsToRealTable(t *testing.T) {
	tests := []struct {
		y, m, d, hh, mm, nights, wantY, wantM, wantD, wantHh, wantMm int
	}{
		{2020, 3, 10, 7, 0, 22, 2020, 3, 11, 9, 24},
		{2020, 3, 10, 7, 0, 65, 2020, 3, 13, 13, 0},
		{2020, 3, 10, 7, 0, 66, 2020, 3, 13, 14, 12},
		{2020, 3, 10, 7, 0, 67, 2020, 3, 13, 15, 24},
	}
	for _, tt := range tests {
		ttTime := time.Date(tt.y, time.Month(tt.m), tt.d, tt.hh, tt.mm, 0, 0, time.Local)
		ttWant := time.Date(tt.wantY, time.Month(tt.wantM), tt.wantD, tt.wantHh, tt.wantMm, 0, 0, time.Local)
		gotTime := getNextWonderTime(ttTime, tt.nights)
		if gotTime != ttWant {
			t.Errorf("rangetest to faild. exect:%v, actual:%v", ttWant, gotTime)
		}
	}
}

/*
	daySec := DaySeconds(time.Now())
	fmt.Println(daySec)
	Adate := SecToAstoltiaTime(180)
	fmt.Println("Astoltia time is ", Adate)
	tomin := AstNightsToMin(Adate, 66)
	fmt.Println(tomin) //4752
	fmt.Println(time.Now().Add(time.Duration(tomin) * time.Minute))

	fmt.Println("Astoltia time is ", Adate)
	fmt.Println("Astoltia time is ", SecToAstoltiaTime(int(daySec)))

	ColossusTime := time.Date(2020, 3, 10, 7, 0, 0, 0, time.Local)
	day, hour, min, sec := AstoltiaNights(ColossusTime)
	fmt.Printf("Colossus time is %dNights %2d:%2d:%2d\n", day, hour, min, sec)
	for i := 0; i < 5; i++ {
		fmt.Println(ColossusTime.Add(time.Duration(AstNightsToMin(ColossusTime, i)) * time.Minute))
	}
	for i := 0; i < 660; i += 66 {
		fmt.Println(ColossusTime.Add(time.Duration(AstNightsToMin(ColossusTime, i)) * time.Minute))
	}

	day, hour, min, sec = AstoltiaNights(time.Now())
	fmt.Printf("Astoltia time is %dNights %2d:%2d:%2d\n", day, hour, min, sec)

*/
