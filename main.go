package main

import (
	"log"

	"github.com/adam-kn1ght/insane-torrent/server"
	"github.com/jpillora/opts"
)

var VERSION = "1.1.0-beta" //set with ldflags

func main() {
	s := server.Server{
		Title:      "Insane - Torrent",
		Port:       3000,
		ConfigPath: "insane-torrent.json",
	}

	o := opts.New(&s)
	o.Version(VERSION)
	o.PkgRepo()
	o.LineWidth = 96
	o.Parse()

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
