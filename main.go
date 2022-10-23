package main

import (
	"github.com/m-oons/headpat-osc/database"
	"github.com/m-oons/headpat-osc/vrc"
)

func main() {
	database.SetupDatabase()
	vrc.SetupOsc()
}
