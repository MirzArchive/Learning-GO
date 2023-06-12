package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

// func sum[T constraints.Ordered](a T, b T) T {
// 	return a + b
// }

func main() {
	file := os.Args[1:]

	cmd := fmt.Sprintf("go run %s", path.Join("exercises", "noval_agung", strings.Join(file, "_"), "main.go"))

	if output, err := executeCmd(cmd); err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(output)
	}
}

func executeCmd(cmd string) (string, error) {
	var output []byte
	var err error

	if runtime.GOOS == "windows" {
		output, err = exec.Command("cmd", "/C", cmd).Output()
	} else {
		output, err = exec.Command("bsdh", "-c", cmd).Output()
	}

	return string(output), err
}
