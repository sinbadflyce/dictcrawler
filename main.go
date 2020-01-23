package main

import (
	"github.com/sinbadflyce/dictcrawler/database"
	"github.com/sinbadflyce/dictcrawler/service"
)

func main() {
	//var c crawling.Crawler
	//c.AtURL = "https://www.ldoceonline.com/dictionary/push"
	//w := c.Run()
	//r := database.Repository{}
	//r.Open("mongodb://localhost:27017")
	//r.Save(w)
	//w := r.Find("push")
	//fmt.Println(w)
	//r.Close()

	database.DictRepo.Open("mongodb://dic.sinbadflyce.com:27017")
	n := service.Network{}
	n.Listen()

	//var c crawling.Crawler
	//c.AtURL = "https://www.ldoceonline.com/dictionary/hello"
	//w := c.Run()
	//database.DictRepo.Save(w)
	database.DictRepo.Close()
}
