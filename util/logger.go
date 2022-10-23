package util

import (
	"fmt"
	"time"
)

func Log(text string, a ...any) {
	ts := timestamp()
	data := fmt.Sprintf(text, a...)
	fmt.Printf("%s %s\n", ts, data)
}

func timestamp() string {
	now := time.Now()
	ts := now.Format("15:04:05")
	return fmt.Sprintf("[%s]", ts)
}
