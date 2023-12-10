package main

import (
	"flag"
	"notesServer/controller/httpserver"
	"notesServer/gates/storage"
	"notesServer/gates/storage/list"
	"notesServer/gates/storage/mp"
)

func main() {
	var st storage.Storage
	var useList bool
	flag.BoolVar(&useList, "list", false, "A boolean flag")
	flag.Parse()

	if useList {
		st = list.NewList()
	} else {
		st = mp.NewMap()
	}
	hs := httpserver.NewHttpServer(":8080", st)
	hs.Start()
}
