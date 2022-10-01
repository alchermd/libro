INSERT INTO books (title, description, cover_photo_url, author, published_at, pages) 
    VALUES($1, $2, $3, $4, $5, $6) 
    RETURNING id