package main

import (
	"lab8/rssparser"
	"log"
)

func main() {
	parser, err := rssparser.NewRSSParser()
	if err != nil {
		log.Fatal(err)
	}
	defer parser.CLoseConnection()
	err = parser.SaveData()
	if err != nil {
		log.Fatal(err)
	}
	parser.ShowData()
	//err = parser.DropTable()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
