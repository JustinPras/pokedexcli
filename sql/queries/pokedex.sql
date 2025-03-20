-- name: CreatePokedexEntry :one
INSERT INTO pokedex(pokemon_name, pokemon_id, json_data)
VALUES (
    ?,
    ?,
    ?
)
RETURNING *;

-- name: GetPokedex :many
SELECT * FROM pokedex;

-- name: GetPokemonByName :one
SELECT * FROM pokedex
WHERE pokemon_name = ?;