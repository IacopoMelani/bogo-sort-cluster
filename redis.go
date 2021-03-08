package main

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if err := rdb.SetNX(ctx, "counter", 0, 0).Err(); err != nil {
		panic(err)
	}
	if err := rdb.Set(ctx, "terminato", false, 0).Err(); err != nil {
		panic(err)
	}
}

func getCounter(ctx context.Context) int {

	mu.Lock()
	defer mu.Unlock()

	val, err := rdb.Get(ctx, "counter").Result()
	if err != nil {
		panic(err)
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	return intVal
}

func incrementCounter(ctx context.Context) {

	mu.Lock()
	defer mu.Unlock()

	if _, err := rdb.Incr(ctx, "counter").Result(); err != nil {
		panic(err)
	}
}

func isTerminato() bool {

	val, err := rdb.Get(ctx, "terminato").Result()
	if err != nil {
		panic(err)
	}

	boolBal, err := strconv.ParseBool(val)
	if err != nil {
		panic(err)
	}

	return boolBal
}

func setTerminato(ctx context.Context) {

	mu.Lock()
	defer mu.Unlock()

	if err := rdb.Set(ctx, "terminato", true, 0).Err(); err != nil {
		panic(err)
	}
}
