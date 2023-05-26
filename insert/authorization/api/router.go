package api

import "net/http"

func Router() *http.ServeMux {
	var mux = http.NewServeMux()

	mux.HandleFunc("/reg", handlerReg)
	mux.HandleFunc("/auth", handlerAuth)
	return mux
}
