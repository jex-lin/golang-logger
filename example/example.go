package main

import (
	"fmt"
	"os"

	"github.com/jex-lin/golang-logger"
)

func main() {
	var log = logger.New(os.Stdout)
	log.SetLevel("warn")
	log.SetTrigger("critical", do)
	log.Notice("Notice") // won't print
	log.Warn("warn")
	log.Error("error")
	log.Critical("critical")
}

func do() {
	fmt.Println("Critical happened.")
}
