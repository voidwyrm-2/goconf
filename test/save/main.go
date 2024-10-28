package main

import (
	"fmt"

	"github.com/voidwyrm-2/goconf"
)

func main() {
	m := map[string]any{
		"name": "flog",
		"type": 0b1010,
		"gen":  10,
		"game": "pocket friends: emblem",
	}

	err := goconf.Save("pf.txt", m)
	if err != nil {
		fmt.Println(err.Error())
	}
}
