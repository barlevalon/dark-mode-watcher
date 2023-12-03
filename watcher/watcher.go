package watcher

import (
	"fmt"
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

	for {
		cmd := exec.Command("/bin/bash", "-c", commandString)
		out, _ := cmd.Output()
		temp := string(out)
		if temp != appearance {
			appearance = temp
			exec.Command(scriptPath, appearance).Start()
		}
		time.Sleep(time.Second * 5)
	}
}
