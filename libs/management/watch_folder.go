package management

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

type FileOp uint32

const  (
	Create FileOp = 1 << iota
	Write
	Remove
	Rename
	Chmod
	Move = Create | Rename
)


func WatchPath(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	log.Println("Watcher created. Watching path: " + path)

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event: ", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file: ", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error", err)
			}
		}
	}()


	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Finished watching: " + path)
	<-done
}
