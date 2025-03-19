-- +goose Up
CREATE TABLE pokedex_abilities(
    id INTEGER PRIMARY KEY,
    abilities_id INTEGER NOT NULL REFERENCES abilities(id) ON DELETE CASCADE,
    pokemon_id INTEGER NOT NULL REFERENCES pokedex(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE pokedex_abilities;