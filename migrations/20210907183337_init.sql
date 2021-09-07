-- +goose Up
-- +goose StatementBegin
CREATE table location
(
    id         SERIAL PRIMARY KEY,
    user_id    INT          NOT NULL,
    address    VARCHAR(255) NOT NULL,
    longitude  FLOAT        NOT NULL,
    latitude   FLOAT        NOT NULL,
    created_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE location             IS 'Содержит данные о местоположении пользователя';
COMMENT ON COLUMN location.user_id    IS 'Пользователь';
COMMENT ON COLUMN location.address    IS 'Адрес';
COMMENT ON COLUMN location.longitude  IS 'Долгота';
COMMENT ON COLUMN location.latitude   IS 'Широта';
COMMENT ON COLUMN location.created_at IS 'Дата и время создания';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE location;
-- +goose StatementEnd
