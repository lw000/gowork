// fsnotify_test project main.go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type Watch struct {
	watch *fsnotify.Watcher
}

func (w *Watch) watchDir(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}

			err = w.watch.Add(path)
			if err != nil {
				return err
			}
			// fmt.Println("monit:", path)
		}

		return nil
	})
	if err != nil {

	}

	go func() {
		for {
			select {
			case ev := <-w.watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						fmt.Println("Create:", ev.Name)

						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Add(ev.Name)
							fmt.Println("add monit:", ev.Name)
						}
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						fmt.Println("Write:", ev.Name)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						fmt.Println("Remove:", ev.Name)

						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Remove(ev.Name)
							fmt.Println("remove monit:", ev.Name)
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						fmt.Println("Rename:", ev.Name)
						w.watch.Remove(ev.Name)
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						fmt.Println("Chmod:", ev.Name)
					}
				}
			case err := <-w.watch.Errors:
				{
					fmt.Println("error:", err)
					return
				}
			}
		}
	}()
}

func main() {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	w := Watch{
		watch: watch,
	}
	w.watchDir("F:\\Project\\Code\\yiwang")

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)

	<-ch

	fmt.Println("exit")

	// select {}
}
