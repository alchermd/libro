CREATE TABLE IF NOT EXISTS books (
    id serial PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL, 
    cover_photo_url TEXT NOT NULL,
    author VARCHAR(255) NOT NULL,
    published_at DATE NOT NULL,
    pages INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS genres (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS books_genres (
    book_id INTEGER NOT NULL,
    genre_id INTEGER NOT NULL,
    PRIMARY KEY (book_id, genre_id),
    FOREIGN KEY (book_id) REFERENCES books (id),
    FOREIGN KEY (genre_id) REFERENCES books (id)
);