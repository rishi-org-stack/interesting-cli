//Package manager  manges data in document of mongodb ie when time period is over it will delete that date and doc
package manager

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"

	dt "github.com/rishi-org-stack/cli/events/data"

	ut "github.com/rishi-org-stack/cli/events/utils"
)

var d = "cli"
var c = "event"
var e dt.Event

func Olddates() {
	_, m, day := time.Now().Date()
	var data []bson.M
	data = e.GetAll(d, c)
	for _, val := range data {
		gd, gm := ut.Parse(val["date"].(string))
		if gm > int(m) {

		}
		if gm == int(m) {
			if gd > day {
			}
			if day > gd {
				e.Name = val["name"].(string)
				e.UID = val["id"].(int32)
				e.Date = val["date"].(string)
				e.Delete(d, c)
			}
		}
		if gm < int(m) {
			e.Name = val["name"].(string)
			e.UID = val["id"].(int32)
			e.Date = val["date"].(string)
			e.Delete(d, c)
		}
	}
}
