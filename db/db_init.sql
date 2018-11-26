CREATE DATABASE books_db owner postgres;
\connect books_db
CREATE TABLE books(
	id_book SERIAL PRIMARY KEY,
	name TEXT NOT NULL);
INSERT INTO books (name) VALUES ('Book1'),('Book2'), ('Book3');

