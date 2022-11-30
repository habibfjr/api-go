package domain

import "movies/dto"

type Movies struct {
	MovieID   string `json:"movieid"`
	MovieName string `json:"moviename"`
}

type MovieRepository interface {
	GetMovies() ([]Movies, error)
	GetMovieByID(string) (*Movies, error)
	InsertMovie(Movies) (*Movies, error)
	DeleteMovie(string) (*Movies, error)
	UpdateMovie(string, Movies) (*Movies, error)
	DeleteAllMovie() (*Movies, error)
}

func (m Movies) ConvertToDTO() dto.MovieResponse {
	return dto.MovieResponse{
		MovieID:   m.MovieID,
		MovieName: m.MovieName,
	}
}
