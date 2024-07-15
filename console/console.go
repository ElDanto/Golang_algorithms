package console

import (
	"os"
	"os/exec"
	"runtime"
	"time"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CloseConsoleWithTimer() {
	time.Sleep(1 * time.Second)
}
