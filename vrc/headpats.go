package vrc

import "github.com/m-oons/headpat-osc/database"

var headpats int64

func SetupHeadpats() {
	headpats = database.GetHeadpatCount()
}

func AddHeadpat() {
	if database.CreateHeadpat() {
		headpats++
	}
}

func GetHeadpats() int64 {
	return headpats
}
