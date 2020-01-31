package main

import (
	"fmt"

	"github.com/sevlyar/go-daemon"
	"github.com/sinbadflyce/dictcrawler/database"
	"github.com/sinbadflyce/dictcrawler/service"
)

func main() {

	context := &daemon.Context{
		PidFileName: "dictcrawler.pid",
		PidFilePerm: 0644,
		LogFileName: "dictcrawler.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon dictcrawler]"},
	}

	d, err := context.Reborn()
	if err != nil {
		fmt.Println("Unable to run: ", err)
	}
	if d != nil {
		return
	}

	defer context.Release()
	runServer()
}

func runServer() {
	database.DictRepo.Open("mongodb://dic.sinbadflyce.com:27017")
	n := service.Network{}
	n.Listen()
	database.DictRepo.Close()
}

func mainDebug() {
	runServer()
}
