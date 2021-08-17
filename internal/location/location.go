package location

import (
	"fmt"
	"time"
)

type Location struct {
	Id        uint64
	UserId    uint64
	Address   string
	Longitude float64
	Latitude  float64
	CreatedAt time.Time
}

func (location Location) String() string {
	return fmt.Sprintf("%s", location.Address)
}
