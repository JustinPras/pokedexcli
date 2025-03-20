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
    pokemon_id INTEGER NOT NULL,
    json_data TEXT NOT NULL
);
