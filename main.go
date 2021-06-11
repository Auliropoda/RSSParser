package main

import (
	"lab8/repository"
	"log"
)

func main() {
	/*if err := parser.Parse(); err != nil {
		log.Fatal(err.Error())
	}*/

	repo, err := repository.NewPostgreRSSRepository()
	if err != nil {
		log.Fatal(err)
	}
	println(repo)
}
