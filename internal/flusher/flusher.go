package flusher

import (
	"github.com/ozonva/ova-location-api/internal/location"
	"github.com/ozonva/ova-location-api/internal/repo"
	"github.com/ozonva/ova-location-api/internal/utils"
)

type Flusher interface {
	Flush(entities []location.Location) []location.Location
}

func NewFlusher(
	chunkSize int,
	entityRepo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

type flusher struct {
	chunkSize  int
	entityRepo repo.Repo
}

func (flusher *flusher) Flush(entities []location.Location) []location.Location {
	chunks := utils.LocationSliceSplit(entities, flusher.chunkSize)
	var unsavedEntities []location.Location
	for _, chunk := range chunks {
		err := flusher.entityRepo.AddEntities(chunk)
		if err != nil {
			unsavedEntities = append(unsavedEntities, chunk...)
		}
	}

	return unsavedEntities
}
