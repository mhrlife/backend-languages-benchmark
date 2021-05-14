package main

import (
	"github.com/labstack/echo"
	"gobench/bucket"
	"math/rand"
	"net/http"
	"runtime"
)

func main() {
	b := bucket.NewBucket()
	runtime.GOMAXPROCS(1)
	e := echo.New()
	e.GET("/get", func(c echo.Context) error {
		resp := ""
		for i := 0; i < 50; i++ {
			if item, founded := b.Get(rand.Intn(1000)); founded {
				resp += item.Value
			}
		}
		return c.String(http.StatusOK, resp)
	})
	e.Logger.Fatal(e.Start(":8080"))

}
