package main

import (
	"fmt"
	"net/http"
)

//goland:noinspection SpellCheckingInspection
func beforeWin10Ver1607(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	//goland:noinspection SpellCheckingInspection
	_, _ = w.Write([]byte("Microsoft NCSI"))
}

//goland:noinspection SpellCheckingInspection
func afterWin10Ver1607(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	//goland:noinspection SpellCheckingInspection
	_, _ = w.Write([]byte("Microsoft Connect Test"))
}

func ping(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			_, _ = fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

//goland:noinspection SpellCheckingInspection
func main() {
	http.HandleFunc("/ncsi.txt", beforeWin10Ver1607)
	http.HandleFunc("/connecttest.txt", afterWin10Ver1607)
	http.HandleFunc("/ping", ping)

	_ = http.ListenAndServe(":80", nil)
}
