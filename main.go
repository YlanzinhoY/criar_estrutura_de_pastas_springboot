package main

import (
	"log"
	"os"
)

func dir() {

	dirs := []string{
		"configuration",
		"controller",
		"exception",
		"dto",
		"repository",
		"services",
	}

	for _, v := range dirs {
		err := os.Mkdir(v, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	err := os.Mkdir("services/impl", os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	dir()
}
