package main

import (
	"log"

	"github.com/maiseea/Bajuse/server"
	"github.com/jpillora/opts"
)

var VERSION = "1.0.0-beta" //set with ldflags

func main() {
	s := server.Server{
		Title:      "Insane - Torrent",
		Port:       3000,
		ConfigPath: "cloud-torrent.json",
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
