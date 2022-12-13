package filewatcher

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

// Watch ...
func Watch(configfile string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("error:", err)
	}
	// errcheck ignore
	defer watcher.Close()

	err = watcher.Add(configfile)
	if err != nil {
		//handle err
		panic(err)
	}
	err = watcher.Add(configfile)
	if err != nil {
		//handle err
		panic(err)
	}
	for {
		select {
		case event := <-watcher.Events:
			// k8s configmaps uses symlinks, we need this workaround.
			// original configmap file is removed
			if event.Op == fsnotify.Remove {
				// remove the watcher since the file is removed
				fmt.Printf("remove: %+v\n", event)
			}
			// also allow normal files to be modified and reloaded.
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Printf("write: %+v\n", event)
			}
			fmt.Printf("received event: %+v\n", event)
		case err := <-watcher.Errors:
			// handle error
			fmt.Printf("err: %+v", err)
			return
		}
	}
}
