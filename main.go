package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	ctx = context.Background()

	mu = sync.Mutex{}

	rdb *redis.Client

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
)

func main() {

	e := echo.New()

	e.Use(middleware.Recover())

	go func() {

		arr = Sort(arr)
		setTerminato(ctx)
	}()

	e.GET("/", func(c echo.Context) error {

		counter := getCounter(c.Request().Context())

		message := fmt.Sprint(counter)

		if isTerminato() {

			message = message + " Finito!"
		}

		return c.HTML(http.StatusOK, message)
	})

	e.Logger.Fatal(e.Start(":8888"))
}
