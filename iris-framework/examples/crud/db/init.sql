DROP TABLE IF EXISTS people;

CREATE TABLE people (
   id SERIAL,
   name TEXT NOT NULL,
   hobby TEXT NOT NULL
);

INSERT INTO people (name, hobby) VALUES ('Johny', 'Jenga');
