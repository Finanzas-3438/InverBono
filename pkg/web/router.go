package web

import "net/http"

func Router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pkg/web/views/index.html")
	})

}
