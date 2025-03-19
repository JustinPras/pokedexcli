CREATE TABLE goose_db_version (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		version_id INTEGER NOT NULL,
		is_applied INTEGER NOT NULL,
		tstamp TIMESTAMP DEFAULT (datetime('now'))
	);
CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE pokedex(
    id INTEGER PRIMARY KEY,
    pokemon_name TEXT NOT NULL,
    experience INTEGER NOT NULL,
    captured_at TIMESTAMP NOT NULL,
    height INTEGER NOT NULL,
    weight INTEGER NOT NULL,
    pokemon_id INTEGER NOT NULL
);
CREATE TABLE abilities(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);
CREATE TABLE pokedex_abilities(
    id INTEGER PRIMARY KEY,
    abilities_id INTEGER NOT NULL REFERENCES abilities(id) ON DELETE CASCADE,
    pokemon_id INTEGER NOT NULL REFERENCES pokedex(id) ON DELETE CASCADE
);
