package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type function func()

func StringNoSpaces(arg string) bool {
	return !strings.Contains(arg, " ")
}

func Cap(arg string) string {
	var capped string

	for i, c := range strings.Split(arg, "") {
		if i == 0 {
			capped += strings.ToUpper(c)
		} else {
			capped += c
		}
	}

	return capped
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error:\t", err.Error())
	}
}

func ExecuteAfterTime(seconds int, f function) {
	duration := time.Duration(seconds) * time.Second
	timer := time.NewTimer(duration)
	<-timer.C
	f()
}

func ExitProg(exitCode int) {
	os.Exit(exitCode)
}

func ComAllArgs() []string {
	return os.Args
}

func ComArgs() []string {
	args := []string{}

	for i, a := range os.Args {
		if i != 0 {
			args = append(args, a)
		}
	}
	return args
}

func ReadStringListToConsole(list []string) {
	for _, l := range list {
		fmt.Println(l)
	}
}
