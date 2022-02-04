package clockfacetest

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"kentliuqiao.com/go-test/clockface"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSvgWriterSecondHand(t *testing.T) {
	cases := []struct {
		t    time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
		{simpleTime(0, 0, 45), Line{150, 150, 60, 150}},
		{simpleTime(0, 0, 15), Line{150, 150, 240, 150}},
	}

	for _, c := range cases {
		t.Run(testName(c.t), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SvgWriter(&b, c.t)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func TestSVGMinuteHand(t *testing.T) {
	cases := []struct {
		t    time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 70}},
	}
	for _, c := range cases {
		t.Run(testName(c.t), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SvgWriter(&b, c.t)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func TestSVGHourHand(t *testing.T) {
	cases := []struct {
		t    time.Time
		line Line
	}{
		{simpleTime(6, 0, 0), Line{150, 150, 150, 200}},
	}
	for _, c := range cases {
		t.Run(testName(c.t), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SvgWriter(&b, c.t)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
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

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if l == line {
			return true
		}
	}
	return false
}
