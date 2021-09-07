package location

import (
	"fmt"
	"time"
)

type Location struct {
	Id        uint64    `db:"id"`
	UserId    uint64    `db:"user_id"`
	Address   string    `db:"address"`
	Longitude float64   `db:"longitude"`
	Latitude  float64   `db:"latitude"`
	CreatedAt time.Time `db:"created_at"`
}

func (location Location) String() string {
	return fmt.Sprintf("%s", location.Address)
}
