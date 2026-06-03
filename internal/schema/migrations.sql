CREATE TABLE IF NOT EXISTS locations (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    image_src TEXT NOT NULL,
    image_alt TEXT NOT NULL,
    lat REAL NOT NULL,
    lon REAL NOT NULL,
    is_dream_location TINYINT NOT NULL DEFAULT 0
);

