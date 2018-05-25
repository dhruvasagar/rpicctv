package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
)

func main() {
	master := gobot.NewMaster()
	a := api.NewAPI(master)
	a.Debug()
	a.Start()

	master.AddRobot(NewCCTVBot())
	master.Start()
}
