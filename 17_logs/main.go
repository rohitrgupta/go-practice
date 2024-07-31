package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
)

func main() {
	var (
		buf bytes.Buffer
	)
	handler := slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(handler)

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	logger.Info("Hello, log file!")
	logger.Warn("hello", "count", 3)

	fmt.Print(&buf)
}
