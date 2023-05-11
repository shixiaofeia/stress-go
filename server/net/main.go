package main

import "net/http"

type Handle struct {
}

func (slf *Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`{"message": "pong"}`))
}

func main() {
	h := new(Handle)
	_ = http.ListenAndServe("0.0.0.0:8080", h)
}
