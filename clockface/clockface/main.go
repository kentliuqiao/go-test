package main

import (
	"os"
	"time"

	"kentliuqiao.com/go-test/clockface"
)

func main() {
	t := time.Now()
	clockface.SvgWriter(os.Stdout, t)
}
