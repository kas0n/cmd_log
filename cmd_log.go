package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {
	shellName := os.Args[1] // "cmd.exe"
	logFileName := os.Args[2]

	c := exec.Command(shellName)
	log_file, _ := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	log_file.WriteString(fmt.Sprintf("-- start log at  %s -- \n", time.Now()))
	defer log_file.Close()

	multiStdout := io.MultiWriter(log_file, os.Stdout)
	multisError := io.MultiWriter(log_file, os.Stderr)
	multiStdin := io.MultiReader(log_file, os.Stdin)

	c.Stdout = multiStdout
	c.Stderr = multisError
	c.Stdin = multiStdin

	c.Run()

}
