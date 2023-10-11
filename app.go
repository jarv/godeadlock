package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/didip/tollbooth/v7"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func main() {
	listenStr := fmt.Sprintf(":%d", 6060)

	lmt := tollbooth.NewLimiter(float64(1), nil)
	lmt.SetOnLimitReached(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Rate limit reached StatusCode: %d\n", lmt.GetStatusCode())
	})

	http.Handle("/", tollbooth.LimitFuncHandler(lmt, testHandler))

	log.Printf("Server started %s\n", listenStr)
	log.Fatal(http.ListenAndServe(listenStr, nil))
}
