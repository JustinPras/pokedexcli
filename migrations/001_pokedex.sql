-- +goose Up
CREATE TABLE pokedex(
    id INTEGER PRIMARY KEY,
    pokemon_name TEXT NOT NULL,
    experience INTEGER NOT NULL,
    captured_at TIMESTAMP NOT NULL,
    height INTEGER NOT NULL,
    weight INTEGER NOT NULL,
    pokemon_id INTEGER NOT NULL
);

-- +goose down
DROP TABLE pokedex;