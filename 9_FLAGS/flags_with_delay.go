package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var name string
var wait time.Duration

var debug bool

// go run .\flags_with_delay.go -name Laura -wait-time 10s

func main() {

	if name == "" {
		fmt.Println("Enter your name:")
		flag.Usage()
		os.Exit(1)
	}

	if debug {
		fmt.Println("Going to wait for ", wait)
	}
	time.Sleep(wait)

	fmt.Println("Hello,", name)
}

func init() {

	flag.BoolVar(&debug, "debug", false, "debug option")
	flag.StringVar(&name, "name", "", "the name to create the hello world message")
	defaultWait, err := time.ParseDuration("5s")

	if err != nil {
		panic("could not parse default wait time")
	}

	flag.DurationVar(&wait, "wait-time", defaultWait, "Time to wait before print")
	flag.Parse()

}
