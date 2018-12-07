CREATE DATABASE books_db owner postgres;

\connect books_db

CREATE TABLE authors(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL);

CREATE TABLE books(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	author_id BIGINT,
	FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE
);


INSERT INTO authors (name) VALUES ('Author1'),('Author2'), ('Author3');
INSERT INTO books (author_id, name) VALUES 	(1, 'Book1_of_Author1'), (1, 'Book2_of_Author1'), (2, 'Book3_of_Author2'), (3, 'Book4_of_Author3');

