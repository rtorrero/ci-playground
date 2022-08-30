package main

import (
	"fmt"
	"moul.io/motd"
	playground "github.com/rtorrero/ci-playground-beta/greetings"
)

func main() {
	fmt.Print(motd.Default())
	fmt.Print(playground.Hello("world"))
}
