CREATE TABLE IF NOT EXISTS users (
    user_id                 UUID          PRIMARY KEY,
    name                    TEXT         NOT NULL,
    age                     INT           NULL
);
