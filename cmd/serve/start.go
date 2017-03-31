package serve

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jaschaephraim/lrserver"
	"gopkg.in/fsnotify.v1"
)

func Start(path string, port int) {

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(path))))
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}
	defer watcher.Close()

	err = watcher.Add(path)
	if err != nil {
		log.Fatalln(err)
	}
	lr := lrserver.New(lrserver.DefaultName, lrserver.DefaultPort)

	go func() {
		log.Fatalln(lr.ListenAndServe())
	}()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println(event)
				lr.Reload(event.Name)
			case err := <-watcher.Errors:
				log.Println(err)
			}
		}
	}()

	log.Fatalln(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}
