package easypdf

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func FileWatcher(files []string, change chan bool) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		dir, _ := os.Getwd()
		for {
			select {
			case event := <-watcher.Events:
				for _, file := range files {
					if file == event.Name || filepath.Join(dir, file) == event.Name {
						change <- true
					}
				}
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()

	for _, file := range files {
		var dir string

		if dir = filepath.Dir(file); dir == "." {
			dir, _ = os.Getwd()
		}

		err = watcher.Add(dir)
		if err != nil {
			log.Fatal(err)
		}
	}

	<-done
}

func DirWatcher(dir string, change chan bool) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-watcher.Events:
				change <- true
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}

	<-done
}
