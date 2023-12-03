package main

import (
	"github.com/barlevalon/dark-mode-watcher/watcher"
)

func main() {
	err := watcher.Watch()
	if err != nil {
		panic(err)
	}
}
