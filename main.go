package main

import (
	"fmt"
	"net/http"
	"os"
)

// newMux 返回所有路由，测试时直接调用它，不用真的监听端口
func newMux(version, host string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "labapp %s running on %s\n", version, host)
	})
	// mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "ok")
	// })
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		fmt.Fprintln(w, "ok")
	})

	return mux
}

func main() {
	host, _ := os.Hostname()
	fmt.Println("listening on :8000")
	http.ListenAndServe(":8000", newMux(os.Getenv("APP_VERSION"), host))
}