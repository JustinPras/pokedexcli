-- name: CreatePokedexEntry :one
INSERT INTO pokedex(pokemon_name, experience, captured_at, height, weight, pokemon_id)
VALUES (
    ?,
    ?,
    ?,
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