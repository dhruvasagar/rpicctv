//go:generate statik -src=./client
package main

import (
	"log"
	"net/http"

	_ "github.com/dhruvasagar/rpicctv/statik"
	"github.com/rakyll/statik/fs"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
)

func main() {
	master := gobot.NewMaster()
	a := api.NewAPI(master)

	fs, err := fs.New()
	if err == nil {
		http.Handle("/cctv/", http.StripPrefix("/cctv/", http.FileServer(fs)))
	} else {
		log.Fatal("There was an error initializing statik fs ", err)
	}

	a.Debug()
	a.Start()

	master.AddRobot(NewCCTVBot())
	master.Start()
}
