package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var num uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		atomic.AddUint64(&num, 1)
		// m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número: %d\n", num)))
	})

	http.ListenAndServe(":3000", nil)
}
