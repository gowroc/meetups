-- +goose Up
CREATE TABLE user_posts (
    user_post_id            UUID          PRIMARY KEY,
    user_id                 UUID          NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    content                 TEXT          NOT NULL
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back


DROP TABLE user_posts;
