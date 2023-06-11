package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

// func sum[T constraints.Ordered](a T, b T) T {
// 	return a + b
// }

func main() {
	file := os.Args[1:2]

	var err error
	var output []byte

	pkg := fmt.Sprintf(
		"go run %s",
		path.Join(
			"exercises",
			"noval_agung",
			file[0],
			"main.go"),
	)

	if runtime.GOOS == "windows" {
		output, err = exec.Command("cmd", "/C", pkg).Output()
	} else {
		output, err = exec.Command("bsdh", "-c", pkg).Output()
	}

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(string(output))
}
