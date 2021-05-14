package main

import (
	"fmt"
	"github.com/labstack/echo"
	"gobench/bucket"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"
)

func main() {
	b := bucket.NewBucket()
	if len(os.Args) > 1 && os.Args[1] == "test" {
		now := time.Now()
		var total int = 1e6
		for i := 0; i < total; i++ {
			b.Get(123)
		}
		done := time.Now().Sub(now).Nanoseconds()
		fmt.Println(done/int64(total), "ns")
	} else {
		runtime.GOMAXPROCS(1)
		e := echo.New()
		e.GET("/get", func(c echo.Context) error {
			resp := ""
			for i := 0; i < 50; i++ {
				if item, _, founded := b.Get(rand.Intn(1000)); founded {
					resp += item.Value
				}
			}
			return c.String(http.StatusOK, resp)
		})
		e.Logger.Fatal(e.Start(":8080"))
	}
}
