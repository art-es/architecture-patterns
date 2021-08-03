package main

import (
	"log"
	"net/http"

	. "github.com/art-es/architecture-patterns/pipe-filter-pattern/common"
)

func main() {
	InitLogger("INITIAL SENDER")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		PublishMessage(ChannelA, "foo")

		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})
	log.Println("[INFO] HTTP server is running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
