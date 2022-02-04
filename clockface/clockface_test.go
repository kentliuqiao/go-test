package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(12, 23, 0), 0},
		{simpleTime(13, 36, 30), math.Pi},
		{simpleTime(15, 45, 15), math.Pi * 15 / 30},
		{simpleTime(12, 12, 7), math.Pi * 7 / 30},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)
			if !roughlyEqualFloat64(got, c.angle) {
				t.Errorf("secondsInRadians() got %v but want %v", got, c.angle)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		t     time.Time
		angle float64
	}{
		{simpleTime(1, 0, 0), 0},
		{simpleTime(1, 30, 0), math.Pi},
		{simpleTime(0, 0, 13), math.Pi * 13 / (30 * 60)},
	}
	for _, c := range cases {
		t.Run(testName(c.t), func(t *testing.T) {
			got := minutesInRadians(c.t)
			if !roughlyEqualFloat64(got, c.angle) {
				t.Errorf("minutesInRadians() got %v but want %v", got, c.angle)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		t     time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), math.Pi * 0},
		{simpleTime(12, 0, 0), math.Pi * 0},
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(3, 0, 0), math.Pi / 2},
		{simpleTime(9, 0, 0), 3 * math.Pi / 2},
		{simpleTime(7, 12, 12), math.Pi / (6 * 3600 / float64(7*3600+12*60+12))},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}
	for _, c := range cases {
		got := hoursInRadians(c.t)
		if !roughlyEqualFloat64(got, c.angle) {
			t.Errorf("hoursInRadians() got %v but want %v", got, c.angle)
		}
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(12, 12, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Errorf("SecondHandVector() got %v but want %v", got, c.point)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		t     time.Time
		point Point
	}{
		{simpleTime(1, 0, 0), Point{0, 1}},
		{simpleTime(1, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run(testName(c.t), func(t *testing.T) {
			got := minuteHandPoint(c.t)
			if !roughlyEqualPoint(got, c.point) {
				t.Errorf("minuteHandPoint() got %v but want %v", got, c.point)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		t     time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run(testName(c.t), func(t *testing.T) {
			got := hourHandPoint(c.t)
			if !roughlyEqualPoint(got, c.point) {
				t.Errorf("hourHandPoint() got %v but want %v", got, c.point)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(1234, time.August, 12, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
