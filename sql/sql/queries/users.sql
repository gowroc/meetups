-- name: SelectAll
SELECT user_id, name, age FROM USERS;

-- name: Insert
INSERT INTO USERS(user_id, name, age) VALUES ($1, $2, $3);

-- name: DeleteAll
DELETE FROM USERS;

-- name: GetByID
SELECT user_id, name, age FROM USERS WHERE user_id = $1;

-- name: GetUserWithPosts
SELECT
u.*, json_agg(up) AS posts
FROM users u
LEFT JOIN user_posts up USING (user_id)
WHERE u.user_id = $1
GROUP BY u.user_id, u.name, u.age;
