package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleep struct {
	Calls int
}

func (s *SpySleep) Sleep() {
	s.Calls++
}

type Op string

const (
	OpSleep = "sleep"
	OpWrite = "write"
)

type SpyOperations struct {
	Calls []Op
	Out   io.Writer
}

func (s *SpyOperations) Sleep() {
	s.Calls = append(s.Calls, OpSleep)
}

func (s *SpyOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, OpWrite)
	return s.Out.Write(p)
}

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintln(out, "Go!")
}

type RealSleep struct{}

func (s RealSleep) Sleep() {
	time.Sleep(time.Second)
}

func main() {
	CountDown(os.Stdout, RealSleep{})
}
