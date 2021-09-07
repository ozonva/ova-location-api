package repo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/ozonva/ova-location-api/internal/location"
)

const table = "location"

type Repo interface {
	AddEntities(entities []location.Location) error
	CreateEntity(ctx context.Context, entity *location.Location) error
	ListEntities(ctx context.Context, limit, offset uint64) ([]location.Location, error)
	GetEntity(ctx context.Context, id uint64) (*location.Location, error)
	RemoveEntity(ctx context.Context, id uint64) (bool, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateEntity(ctx context.Context, entity *location.Location) error {
	query := squirrel.Insert(table).
		Columns("user_id", "address", "longitude", "latitude", "created_at").
		Values(entity.UserId, entity.Address, entity.Longitude, entity.Latitude, entity.CreatedAt).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	err := query.QueryRowContext(ctx).Scan(&entity.Id)

	if err != nil {
		log.Error().Err(err).Msg("Не удалось сохранить локацию")
		return err
	}

	return nil
}

func (r *repo) AddEntities(entities []location.Location) error {
	return nil
}

func (r *repo) GetEntity(ctx context.Context, id uint64) (*location.Location, error) {
	qsql, args, err := squirrel.Select("id", "user_id", "address", "longitude", "latitude", "created_at").
		From(table).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		log.Error().Err(err).Msg("Не удалось сгенерировать запрос на получение локации")
		return nil, err
	}

	var entity location.Location
	err = r.db.Get(&entity, qsql, args...)

	if err != nil {
		log.Error().Err(err).Msg("Не удалось выполнить запрос на получение локации")
		return nil, err
	}

	return &entity, nil
}

func (r *repo) ListEntities(ctx context.Context, limit, offset uint64) ([]location.Location, error) {
	qsql, _, err := squirrel.Select("id", "user_id", "address", "longitude", "latitude", "created_at").
		From(table).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		log.Error().Err(err).Msg("Не удалось сгенерировать запрос на получение списка локаций")
		return nil, err
	}

	var entities []location.Location
	err = r.db.Select(&entities, qsql)
	if err != nil {
		log.Error().Err(err).Msg("Не удалось выполнить запрос на получение списка локаций")
	}

	return entities, nil
}

func (r *repo) RemoveEntity(ctx context.Context, id uint64) (bool, error) {
	qsql, args, err := squirrel.Delete(table).
		Where(squirrel.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		log.Error().Err(err).Msg("Не удалось сгенерировать запрос на удаление локации")
		return false, err
	}

	result, err := r.db.ExecContext(ctx, qsql, args...)
	if err != nil {
		log.Error().Err(err).Msg("Не удалось выполнить запрос на удаление локации")
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected <= 0 {
		log.Error().Err(err).Msg("Локации не существовало")
		return false, errors.New("локации не существовало")
	}

	return true, err
}
