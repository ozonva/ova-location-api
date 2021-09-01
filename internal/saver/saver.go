package saver

import (
	"github.com/ozonva/ova-location-api/internal/flusher"
	"github.com/ozonva/ova-location-api/internal/location"
	"time"
)

type Saver interface {
	Save(entity location.Location)
	Init()
	Close()
	GetBuffer() []location.Location
}

func New(
	capacity uint64,
	flusher flusher.Flusher,
	timeout time.Duration,
) Saver {
	return &saver{
		capacity: capacity,
		flusher:  flusher,
		timeout:  timeout,
		buffer:   make([]location.Location, 0, capacity),
	}
}

type saver struct {
	capacity uint64
	buffer   []location.Location
	flusher  flusher.Flusher
	timeout  time.Duration
	ticker   *time.Ticker
	close    chan bool
	write    chan location.Location
}

func (s *saver) Init() {
	s.ticker = time.NewTicker(s.timeout)
	s.close = make(chan bool)
	s.write = make(chan location.Location)
	s.loop()
}

func (s *saver) Save(entity location.Location) {
	s.write <- entity
}

func (s *saver) GetBuffer() []location.Location {
	return s.buffer
}

func (s *saver) Close() {
	s.close <- true
}

func (s *saver) loop() {
	go func() {
		defer func() {
			s.ticker.Stop()
			close(s.close)
			close(s.write)
		}()
		for {
			select {
			case entity := <-s.write:
				s.save(entity)
			case <-s.close:
				s.flush()
				return
			case <-s.ticker.C:
				s.flush()
			}
		}
	}()
}

func (s *saver) flush() {
	if s.buffer == nil || len(s.buffer) == 0 {
		return
	}

	s.buffer = s.flusher.Flush(s.buffer)
}

func (s *saver) save(entity location.Location) {
	s.buffer = append(s.buffer, entity)
}
