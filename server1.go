package main
require github.com/suwitna/gopl/ch1
import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	lissa(w)
}
