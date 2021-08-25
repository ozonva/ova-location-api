package repo

import "github.com/ozonva/ova-location-api/internal/location"

type Repo interface {
	AddEntities(entities []location.Location) error
	ListEntities(limit, offset uint64) ([]location.Location, error)
	GetEntity(entityId uint64) (*location.Location, error)
}
