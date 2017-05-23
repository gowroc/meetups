-- +goose Up
CREATE TABLE users (
    user_id                 UUID          PRIMARY KEY,
    name                    TEXT          NOT NULL,
    age                     INT           NULL
);

-- +goose Down
DROP TABLE users;
