-- +goose Up
CREATE TABLE abilities(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
)

-- +goose Down
DROP TABLE abilities;