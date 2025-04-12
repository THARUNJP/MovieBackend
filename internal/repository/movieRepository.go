package repository

import (
	"MovieBack/internal/types"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MovieRepository interface {
	GetMovies(ctx context.Context) ([]types.MovieStruct, error)
}

type movieRepo struct {
	db *pgxpool.Pool
}

func NewMovieRepository(db *pgxpool.Pool) MovieRepository {
	return &movieRepo{db: db}
}
func (r *movieRepo) GetMovies(ctx context.Context) ([]types.MovieStruct, error) {
	rows, err := r.db.Query(ctx, "SELECT movie_id, movie_name, movie_genre, is_active FROM movies")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var movies []types.MovieStruct

	for rows.Next() {
		var m types.MovieStruct
		err := rows.Scan(&m.MovieID, &m.MovieName, &m.MovieGenre, &m.IsActive)
		if err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}
	return movies, nil
}

// https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
