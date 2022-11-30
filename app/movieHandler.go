package app

import (
	"encoding/json"
	"fmt"
	"log"
	"movies/dto"
	"movies/helper"
	"movies/service"
	"net/http"

	"github.com/gorilla/mux"
)

type MovieHandler struct {
	service service.MovieService
}

func (mh *MovieHandler) getAllMovies(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		log.Fatalln("wrong method", r.Method)
	} else {
		movies, err := mh.service.GetAll()
		if err != nil {
			response := helper.WriteErrResponse(http.StatusNotFound, "no data in database", helper.EmptyData{})
			log.Fatalln(response)
		}
		fmt.Println("fetching data...")
		response := helper.WriteResponse(http.StatusOK, "data successfully fetched", movies)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func (mh *MovieHandler) getMovieByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movieID := params["movieid"]

	if r.Method != "GET" {
		log.Fatalln("wrong method", r.Method)
	} else {
		movie, err := mh.service.GetByID(movieID)
		if err != nil {
			response := helper.WriteErrResponse(http.StatusBadRequest, "no movie id specified", helper.EmptyData{})
			log.Fatalln(response)
		}
		fmt.Println("fetching data by id...")
		response := helper.WriteResponse(http.StatusOK, "data successfully fetched", movie)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}

func (mh *MovieHandler) addMovie(w http.ResponseWriter, r *http.Request) {
	var input dto.NewMovie

	if r.Method != "POST" {
		log.Fatalln("wrong method", r.Method)
	} else {
		input.MovieID = r.FormValue("movieid")
		input.MovieName = r.FormValue("moviename")
		if input.MovieID == "" || input.MovieName == "" {
			response := helper.WriteErrResponse(http.StatusBadRequest, "one or some parameters are empty", helper.EmptyData{})
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			fmt.Println("one or some parameters are empty")
			return
		}
		movie, err := mh.service.InsertData(input)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("inserting data...")
		response := helper.WriteResponse(http.StatusOK, "data successfully added", movie)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}

func (mh *MovieHandler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movieID := params["movieid"]

	if movieID == "" {
		response := helper.WriteResponse(http.StatusBadRequest, "no movie id specified", helper.EmptyData{})
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		movie, err := mh.service.DeleteData(movieID)
		if err != nil {
			log.Fatalln("failed to delete data", err)
		}
		response := helper.WriteResponse(http.StatusOK, "data has been deleted", movie)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func (mh *MovieHandler) updateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movieID := params["movieid"]

	var input dto.NewMovie

	if movieID == "" {
		response := helper.WriteErrResponse(http.StatusBadRequest, "no movie id specified", helper.EmptyData{})
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		input.MovieName = r.FormValue("moviename")
		movie, err := mh.service.UpdateData(movieID, input)
		if err != nil {
			log.Fatalln("failed to update data")
		}
		fmt.Println("updating data...")
		response := helper.WriteResponse(http.StatusOK, "data successfully updated", movie)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func (mh *MovieHandler) deleteAllMovie(w http.ResponseWriter, r *http.Request) {
	movie, err := mh.service.DeleteAll()
	if err != nil {
		log.Fatalln("failed to delete data", err)
	}
	response := helper.WriteResponse(http.StatusOK, "data has been deleted", movie)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
