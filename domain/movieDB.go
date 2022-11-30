package domain

import (
	"database/sql"
	"fmt"
	"log"
	"movies/helper"
	"net/http"
)

type MovieDB struct {
	db *sql.DB
}

func NewMovieDB(client *sql.DB) MovieDB {
	return MovieDB{client}
}

func (mdb MovieDB) GetMovies() ([]Movies, error) {
	rows, err := mdb.db.Query("select * from movies")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var movies []Movies

	for rows.Next() {
		var id int
		var movieID string
		var movieName string

		err = rows.Scan(&id, &movieID, &movieName)

		if err != nil {
			return nil, err
		}

		movies = append(movies, Movies{
			MovieID:   movieID,
			MovieName: movieName})

	}
	return movies, nil
}

func (mdb MovieDB) GetMovieByID(id string) (*Movies, error) {
	var movieID string
	var movieName string
	var movies []Movies

	row, err := mdb.db.Query("select * from movies where movieid = $1", id)
	if err != nil {
		response := helper.WriteErrResponse(http.StatusBadRequest, "request denied (bad query)", helper.EmptyData{})
		log.Fatalln(response)
	}

	for row.Next() {
		var id int

		err = row.Scan(&id, &movieID, &movieName)
		if err != nil {
			response := helper.WriteErrResponse(http.StatusNotFound, "no data in database", helper.EmptyData{})
			log.Fatalln(response)
		}
	}

	movies = append(movies, Movies{
		MovieID:   movieID,
		MovieName: movieName,
	})
	return &movies[0], nil
}

func (mdb MovieDB) InsertMovie(m Movies) (*Movies, error) {

	var movies []Movies
	var lastInsertID int

	err := mdb.db.QueryRow("insert into movies(movieID, movieName) values($1, $2) returning id;", m.MovieID, m.MovieName).Scan(&lastInsertID)
	if err != nil {
		response := helper.WriteErrResponse(http.StatusBadRequest, "request denied (bad query)", helper.EmptyData{})
		log.Fatalln(response)
	}

	movies = append(movies, Movies{
		MovieID:   m.MovieID,
		MovieName: m.MovieName})

	return &movies[0], nil
}

func (mdb MovieDB) DeleteMovie(id string) (*Movies, error) {
	var movies Movies
	_, err := mdb.db.Exec("delete from movies where movieid = $1", id)
	if err != nil {
		helper.WriteErrResponse(http.StatusBadRequest, "request to delete denied - bad query", helper.EmptyData{})
	}

	fmt.Println("deleted a movie with id: " + id)

	return &movies, nil
}

func (mdb MovieDB) UpdateMovie(id string, m Movies) (*Movies, error) {
	query := "update movies set moviename = $1 where movieid = $2;"
	_, err := mdb.db.Exec(query, m.MovieName, id)
	if err != nil {
		helper.WriteErrResponse(http.StatusBadRequest, "request to update data failed", helper.EmptyData{})
	}

	fmt.Println("updated a movie with id: " + id)

	return &m, nil
}

func (mdb MovieDB) DeleteAllMovie() (*Movies, error) {
	var movies Movies
	_, err := mdb.db.Exec("delete from movies")
	if err != nil {
		helper.WriteErrResponse(http.StatusBadRequest, "request to delete all denied - bad query", helper.EmptyData{})
	}

	fmt.Println("deleted all movies")

	return &movies, nil
}
