CREATE TABLE IF NOT EXISTS person (
    personid SERIAL PRIMARY KEY,
    name     VARCHAR(255) NOT NULL,
    age      INT NOT NULL,
    address  TEXT,
    work     TEXT
);
