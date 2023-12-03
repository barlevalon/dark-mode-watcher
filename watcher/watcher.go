package watcher

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

func Watch() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	commandString := fmt.Sprintf("defaults read %s/Library/Preferences/.GlobalPreferences.plist AppleInterfaceStyle", homeDir)
	appearance := "init"
	scriptPath := path.Join(homeDir, ".config/fish/functions/update_dark_mode.fish")
	logFile, err := os.OpenFile(path.Join(homeDir, "dark-mode-watcher.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	for {
		cmd := exec.Command("/bin/bash", "-c", commandString)
		out, _ := cmd.CombinedOutput()
		log.Println(string(out))
		temp := string(out)
		if temp != appearance {
			appearance = temp
			out, _ := exec.Command(scriptPath, appearance).CombinedOutput()
			log.Println(string(out))
		}
		time.Sleep(time.Second * 5)
	}
}
