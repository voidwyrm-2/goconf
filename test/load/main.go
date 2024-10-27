package main

import (
	"fmt"

	"github.com/voidwyrm-2/goconf"
)

func main() {
	config, err := goconf.Load("config.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// config is a string -> any map

	fmt.Println(config["name"]) // "Jacob Thaumiel"
	fmt.Println(config["age"])  // 21
}
