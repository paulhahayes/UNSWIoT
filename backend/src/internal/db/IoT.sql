CREATE TABLE IF NOT EXISTS devices (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

DELETE FROM devices;

INSERT INTO devices (name) VALUES ('tp-link');
INSERT INTO devices (name) VALUES ('hikvision');