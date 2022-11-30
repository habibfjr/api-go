package service

import (
	"movies/domain"
	"movies/dto"
)

type MovieService interface {
	GetAll() ([]dto.MovieResponse, error)
	GetByID(string) (*dto.MovieResponse, error)
	InsertData(dto.NewMovie) (*dto.MovieResponse, error)
	DeleteData(string) (*dto.MovieResponse, error)
	UpdateData(string, dto.NewMovie) (*dto.MovieResponse, error)
	DeleteAll() (*dto.MovieResponse, error)
}

type DefMovieService struct {
	repo domain.MovieRepository
}

func NewMovieService(repo domain.MovieRepository) DefMovieService {
	return DefMovieService{repo}
}

func (s DefMovieService) GetAll() ([]dto.MovieResponse, error) {
	movies, err := s.repo.GetMovies()
	if err != nil {
		return nil, err
	}

	var res []dto.MovieResponse
	for _, data := range movies {
		res = append(res, data.ConvertToDTO())
	}
	return res, nil
}

func (s DefMovieService) GetByID(id string) (*dto.MovieResponse, error) {
	movie, err := s.repo.GetMovieByID(id)
	if err != nil {
		return nil, err
	}
	// var res dto.MovieResponse
	res := movie.ConvertToDTO()

	return &res, nil
}

func (s DefMovieService) InsertData(nm dto.NewMovie) (*dto.MovieResponse, error) {
	var m domain.Movies

	m.MovieID = nm.MovieID
	m.MovieName = nm.MovieName

	movie, err := s.repo.InsertMovie(m)
	if err != nil {
		return nil, err
	}

	res := movie.ConvertToDTO()

	return &res, nil
}

func (s DefMovieService) DeleteData(id string) (*dto.MovieResponse, error) {

	movie, err := s.repo.DeleteMovie(id)
	if err != nil {
		return nil, err
	}

	var res = movie.ConvertToDTO()

	return &res, nil
}

func (s DefMovieService) UpdateData(id string, um dto.NewMovie) (*dto.MovieResponse, error) {
	var m domain.Movies

	m.MovieID = id
	m.MovieName = um.MovieName

	movie, err := s.repo.UpdateMovie(id, m)
	if err != nil {
		return nil, err
	}

	var res = movie.ConvertToDTO()

	return &res, nil
}

func (s DefMovieService) DeleteAll() (*dto.MovieResponse, error) {

	movie, err := s.repo.DeleteAllMovie()
	if err != nil {
		return nil, err
	}

	var res = movie.ConvertToDTO()

	return &res, nil
}
