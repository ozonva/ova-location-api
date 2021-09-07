package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ozonva/ova-location-api/internal/location"
	"github.com/ozonva/ova-location-api/internal/repo"
	api "github.com/ozonva/ova-location-api/pkg/ova-location-api"
)

type OvaLocationApi struct {
	api.UnimplementedApiServer
	repo repo.Repo
}

func New(repo repo.Repo) api.ApiServer {
	return &OvaLocationApi{
		repo: repo,
	}
}

func (a *OvaLocationApi) CreateLocationV1(ctx context.Context, req *api.CreateLocationV1Request) (*api.LocationV1Response, error) {
	entity := &location.Location{
		UserId:    req.UserId,
		Address:   req.Address,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	}

	err := a.repo.CreateEntity(ctx, entity)

	if err != nil {
		log.Error().Err(err).Msg("Не удалось сохранить локацию")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.LocationV1Response{
		Id:        entity.Id,
		UserId:    entity.UserId,
		Address:   entity.Address,
		Longitude: entity.Longitude,
		Latitude:  entity.Latitude,
		CreatedAt: timestamppb.New(entity.CreatedAt),
	}, nil
}

func (a *OvaLocationApi) GetLocationV1(ctx context.Context, req *api.GetLocationV1Request) (*api.LocationV1Response, error) {
	entity, err := a.repo.GetEntity(ctx, req.Id)

	if err != nil {
		log.Error().Err(err).Msg("Локация не найдена")
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &api.LocationV1Response{
		Id:        entity.Id,
		UserId:    entity.UserId,
		Address:   entity.Address,
		Longitude: entity.Longitude,
		Latitude:  entity.Latitude,
		CreatedAt: timestamppb.New(entity.CreatedAt),
	}, nil
}

func (a *OvaLocationApi) ListLocationsV1(ctx context.Context, req *api.ListLocationV1Request) (*api.ListLocationsV1Response, error) {
	entities, err := a.repo.ListEntities(ctx, req.Limit, req.Offset)

	if err != nil {
		log.Error().Err(err).Msg("Не удалось получить запрашиваемые данные")
		return nil, status.Error(codes.Internal, err.Error())
	}

	locations := make([]*api.LocationV1Response, 0, len(entities))

	for _, entity := range entities {
		locations = append(locations, &api.LocationV1Response{
			Id:        entity.Id,
			UserId:    entity.UserId,
			Address:   entity.Address,
			Longitude: entity.Longitude,
			Latitude:  entity.Latitude,
			CreatedAt: timestamppb.New(entity.CreatedAt),
		})
	}

	return &api.ListLocationsV1Response{
		List: locations,
	}, nil
}

func (a *OvaLocationApi) RemoveLocationV1(ctx context.Context, req *api.RemoveLocationV1Request) (*api.RemoveV1Response, error) {
	removed, err := a.repo.RemoveEntity(ctx, req.Id)

	if err != nil {
		log.Error().Err(err).Msg("Не удалось удалить локацию")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.RemoveV1Response{
		Status: removed,
	}, nil
}
