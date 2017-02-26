package serve

import (
	"fmt"
	"log"
	"net/http"
)

func Start(path string, port int) {

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(path))))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}
