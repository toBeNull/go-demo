package main

import (
	"fmt"
	"go-hello/version"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := ":8080"
	startTime := time.Now()
	appVersion := version.AppVersion()
	format := "2006.01.02 15:04:05"

	nodeName, _ := os.LookupEnv("MY_NODE_NAME")
	nodeIp, _ := os.LookupEnv("MY_NODE_IP")
	podName, _ := os.LookupEnv("MY_POD_NAME")
	podIp, _ := os.LookupEnv("MY_POD_IP")
	namespace, _ := os.LookupEnv("MY_POD_NAMESPACE")
	serviceAccountName, _ := os.LookupEnv("MY_POD_SERVICE_ACCOUNT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%v] receive request... \n", time.Now().Format(format))

		fmt.Fprintln(w, "Welcome to website!")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "version    : ", appVersion)
		fmt.Fprintln(w, "startTime  : ", startTime.Format(format))
		fmt.Fprintln(w, "currentTime: ", time.Now().Format(format))
		fmt.Fprintln(w, "")

		fmt.Fprintln(w, "host name  : ", nodeName)
		fmt.Fprintln(w, "host ip    : ", nodeIp)

		fmt.Fprintln(w, "pod name   : ", podName)
		fmt.Fprintln(w, "pod ip     : ", podIp)

		fmt.Fprintln(w, "namespace  : ", namespace)
		fmt.Fprintln(w, "service account : ", serviceAccountName)

		fmt.Fprintln(w, r.URL.Query().Get("name"))
	})

	fmt.Println("server will be ready to handle request...")
	fmt.Println("listen on: ", addr)
	http.ListenAndServe(addr, nil)
}
