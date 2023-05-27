package seed

import (
	ei "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func CreateStatus() []*ei.Status {
	status := []*ei.Status{
		{
			StatusInformation: "Terbit",
		},
		{
			StatusInformation: "Draft",
		},
	}
	return status
}
