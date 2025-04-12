package container

import (
	"MovieBack/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Container holds all application dependencies
type Container struct {
	MovieRepo repository.MovieRepository
	
}

// NewContainer creates a new container with initialized dependencies
func NewContainer(db *pgxpool.Pool) *Container {
	return &Container{
		MovieRepo: repository.NewMovieRepository(db),
	}
}
