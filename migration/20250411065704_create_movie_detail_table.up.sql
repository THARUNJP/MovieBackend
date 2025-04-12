CREATE TABLE IF NOT EXISTS movie_details (
    detail_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    movie_id UUID REFERENCES movies(movie_id),
    description TEXT NOT NULL,
    movie_cast JSONB NOT NULL,
    language VARCHAR(100) NOT NULL,
    release_date DATE NOT NULL
);
