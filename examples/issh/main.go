// MIT License Copyright (C) 2022 Hiroshi Shimamoto
package main

import (
	"log"
	"os"

	"github.com/hshimamoto/go-issh"
)

func main() {
	if len(os.Args) != 4 {
		log.Printf("issh addr user keyfile")
		return
	}

	addr := os.Args[1]
	user := os.Args[2]
	keyfile := os.Args[3]

	config, err := issh.NewConfig(user, keyfile)
	if err != nil {
		log.Printf("Config: %v", err)
		return
	}
	cli, err := issh.Dial(addr, config)
	if err != nil {
		log.Printf("Dial: %v", err)
		return
	}
	defer cli.Close()

	session, err := cli.NewSession()
	if err != nil {
		log.Printf("NewSession: %v", err)
		return
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	err = session.Shell()
	if err != nil {
		log.Printf("Shell: %v", err)
		return
	}
}
