CREATE DATABASE readingList;
CREATE ROLE readinglist WITH LOGIN PASSWORD 'password';

CREATE TABLE IF NOT EXISTS books (

    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    published integer NOT NULL,
    pages integer NOT NULL,
    rating real NOT NULL

);

GRANT SELECT,INSERT,UPDATE, DELETE ON books TO readinglist;
GRANT USAGE, SELECT ON SEQUENCE  books_id_seq TO readinglist;