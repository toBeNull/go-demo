package main

import (
	"fmt"
	"go-hello/version"
	"net/http"
	"time"
)

func main() {
	startTime := time.Now()
	appVersion := version.AppVersion()
	format := "2006.01.02 15:04:05"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%v] receive request... \n", time.Now().Format(format))

		fmt.Fprintln(w, "Welcome to website!")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "version    : ", appVersion)
		fmt.Fprintln(w, "startTime  : ", startTime.Format(format))
		fmt.Fprintln(w, "currentTime: ", time.Now().Format(format))
		fmt.Fprintln(w, r.URL.Query().Get("name"))
	})

	fmt.Println("server will be ready to handle request...")
	http.ListenAndServe(":8080", nil)
}
