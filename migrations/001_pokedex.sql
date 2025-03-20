-- +goose Up
CREATE TABLE pokedex(
    id INTEGER PRIMARY KEY,
    pokemon_name TEXT NOT NULL,
    pokemon_id INTEGER NOT NULL,
    json_data TEXT NOT NULL
);

-- +goose down
DROP TABLE pokedex;