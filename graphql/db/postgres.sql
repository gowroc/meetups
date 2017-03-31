CREATE DATABASE graphql;

\c graphql;

CREATE TABLE users (
    username    varchar(30) CONSTRAINT firstkey PRIMARY KEY,
    admin       boolean
);

INSERT INTO users(username, admin) VALUES ('slomek',true);
