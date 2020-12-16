CREATE TABLE books (
  id serial,
  title varchar,
  author varchar,
  year varchar
);

INSERT INTO books(title, author, year) VALUES
  ('Golang is great', 'Mr. Great', '2012'),
  ('Golang isnt bad', 'Mr. Bad', '2013');