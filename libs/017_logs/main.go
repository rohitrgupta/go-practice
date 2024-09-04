package main

import (
	"bytes"
	"context"
	"log"
	"log/slog"
	"strings"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	var (
		buf bytes.Buffer
	)
	handler := slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true})
	logger := slog.New(handler)

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	logger.Info("Hello, log file!")
	logger.Warn("hello", "count", 3)
	logger.Debug("hello", "count", 3)

	logger.Error("hello", "count", 3)
	logger.WithGroup("decoder").Info("hello", "count", 3)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       15, // use default DB
	})

	// values := make([]strinng, 0)
	v := strings.Split(buf.String(), "\n")
	// fmt.Println(v)
	var v1 []interface{}
	for _, v := range v {
		v1 = append(v1, v)
	}
	err := rdb.RPush(ctx, "key1", v1...).Err()
	if err != nil {
		panic(err)
	}
	// fmt.Print(&buf)
}
