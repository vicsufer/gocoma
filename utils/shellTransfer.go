package utils

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func SwitchShell(path string) {

	usr, _ := user.Current()
	dir := usr.HomeDir

	if path == "~" {
		// In case of "~", which won't be caught by the "else if"
		path = dir
	} else if strings.HasPrefix(path, "~/") || strings.HasPrefix(path, "$HOME"){
		// Use strings.HasPrefix so we don't match paths like
		// "/something/~/something/"
		path = filepath.Join(dir, path[2:])
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("The configured repo directory does not exist: %s\n %s", dir, path)
		os.Exit(1)
	}

	me, err := user.Current()
	if err != nil {
		panic(err)
	}

	shellPath := os.Getenv("SHELL")
	shellPathArray := strings.Split(shellPath, "/")
	shell := shellPathArray[len(shellPathArray)-1]

	var args = make([]string, 2)
	switch shell {
	case "zsh":
		args = []string{""}
	default:
		args = []string{"-fpl", me.Username}
	}

	// Transfer stdin, stdout, and stderr to the new process
	// and also set target directory for the shell to start in.
	pa := os.ProcAttr {
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir: path,
	}

	proc, err := os.StartProcess(shellPath, args, &pa)
	if err != nil {
		panic(err)
	}

	// Wait until user exits the shell
	_, err = proc.Wait()
	if err != nil {
		panic(err)
	}
}