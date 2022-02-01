package main

import (
	"log"
	"yagc/cmd"
)

func main() {
	log.SetFlags(log.Lmsgprefix)
	log.SetPrefix("[yagc] ")
	cmd.Execute()
}
