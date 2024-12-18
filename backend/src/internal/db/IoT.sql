CREATE TABLE IF NOT EXISTS devices (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

DELETE FROM devices;

INSERT INTO devices (name) VALUES ('test_device_1');
INSERT INTO devices (name) VALUES ('test_device_2');