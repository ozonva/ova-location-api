package api

import (
	"context"
	api "github.com/ozonva/ova-location-api/pkg/ova-location-api"
	"github.com/rs/zerolog/log"
)

type OvaLocationApi struct {
	api.UnimplementedApiServer
}

func New() api.ApiServer {
	return &OvaLocationApi{}
}

func (a *OvaLocationApi) CreateLocationV1(ctx context.Context, req *api.CreateLocationV1Request) (*api.LocationV1Response, error) {
	log.Info().Msgf("CreateLocationV1 request: %v", req)
	return a.UnimplementedApiServer.CreateLocationV1(ctx, req)
}

func (a *OvaLocationApi) GetLocationV1(ctx context.Context, req *api.GetLocationV1Request) (*api.LocationV1Response, error) {
	log.Info().Msgf("GetLocationV1 request: %v", req)
	return a.UnimplementedApiServer.GetLocationV1(ctx, req)
}

func (a *OvaLocationApi) ListLocationsV1(ctx context.Context, req *api.ListLocationV1Request) (*api.ListLocationsV1Response, error) {
	log.Info().Msgf("ListLocationsV1 request: %v", req)
	return a.UnimplementedApiServer.ListLocationsV1(ctx, req)
}

func (a *OvaLocationApi) RemoveLocationV1(ctx context.Context, req *api.RemoveLocationV1Request) (*api.RemoveV1Response, error) {
	log.Info().Msgf("RemoveLocationV1 request: %v", req)
	return a.UnimplementedApiServer.RemoveLocationV1(ctx, req)
}
