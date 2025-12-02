package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"Kate.com/TIP_pr13/internal/work"
)

func main() {

	runtime.SetBlockProfileRate(1)     // профиль блокировок
	runtime.SetMutexProfileFraction(1) // профиль мьютексов

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		defer work.TimeIt("Fib(38)")()

		n := 38
		res := work.FibFast(n)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte(fmtInt(res)))
	})

	log.Println("Server on :8080; pprof on /debug/pprof/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fmtInt(v int) string { return fmt.Sprintf("%d\n", v) }
