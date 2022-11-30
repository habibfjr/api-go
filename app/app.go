package app

import (
	"database/sql"
	"fmt"
	"log"
	"movies/domain"
	"movies/service"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}

	for _, envKey := range envProps {
		if os.Getenv(envKey) == "" {
			log.Fatal(fmt.Printf("environment variable %s not set. terminating application...", envKey))
		}
	}

	log.Println("environment variables loaded...")
}

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	log.Println("load env variables")

	sanityCheck()

	dbClient := getClientDB()
	// fmt.Println(dbClient)

	movieRepositoryDB := domain.NewMovieDB(dbClient)

	movieService := service.NewMovieService(movieRepositoryDB)

	mh := MovieHandler{movieService}

	router := mux.NewRouter()

	router.HandleFunc("/movies", mh.getAllMovies).Methods("GET")

	router.HandleFunc("/movies/{movieid}", mh.getMovieByID).Methods("GET")

	router.HandleFunc("/movies", mh.addMovie).Methods("POST")

	router.HandleFunc("/movies/{movieid}", mh.deleteMovie).Methods("DELETE")

	router.HandleFunc("/movies", mh.deleteAllMovie).Methods("DELETE")

	router.HandleFunc("/movies/{movieid}", mh.updateMovie).Methods("PUT")

	fmt.Println("Server start at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getClientDB() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbName := os.Getenv("DB_NAME")

	source := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPasswd, dbName)

	db, err := sql.Open("postgres", source)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

// type Movies struct {
// 	MovieID   string `json:"movieid"`
// 	MovieName string `json:"moviename"`
// }

// type JsonResponse struct {
// 	Status  int         `json:"status"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

// type EmptyData struct{}

// func writeResponse(status int, message string, data interface{}) JsonResponse {
// 	return JsonResponse{
// 		Status:  status,
// 		Message: message,
// 		Data:    data,
// 	}
// }

// func writeErrResponse(status int, message string, data interface{}) JsonResponse {
// 	// split := strings.Split(err, "\n")
// 	return JsonResponse{
// 		Status:  status,
// 		Message: message,
// 		Data:    data,
// 	}
// }

// func GetMovies(w http.ResponseWriter, r *http.Request) {
// 	db := getClientDB()

// 	fmt.Println("getting data")
// 	rows, err := db.Query("select * from movies")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var movies []Movies

// 	for rows.Next() {
// 		var id int
// 		var movieID string
// 		var movieName string

// 		err = rows.Scan(&id, &movieID, &movieName)

// 		if err != nil {
// 			response := writeErrResponse(http.StatusNotFound, "no data in database", EmptyData{})
// 			log.Fatalln(response)
// 		}

// 		movies = append(movies, Movies{
// 			MovieID:   movieID,
// 			MovieName: movieName})

// 	}
// 	// var response = JsonResponse{Status: http.StatusOK, Data: movies, Message: "data successfully fetched"}
// 	response := writeResponse(http.StatusOK, "data successfully fetched", movies)

// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func InsertMovie(w http.ResponseWriter, r *http.Request) {
// 	movieID := r.FormValue("movieid")
// 	movieName := r.FormValue("moviename")

// 	var movies []Movies

// 	var response = JsonResponse{}

// 	if movieID == "" || movieName == "" {
// 		response := writeErrResponse(http.StatusBadRequest, "no value added to parameters", EmptyData{})
// 		log.Println(response)
// 		w.Header().Add("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(response)
// 	} else {
// 		db := getClientDB()

// 		fmt.Println("adding a movie with id: " + movieID + " and name: " + movieName)

// 		var lastInsertID int

// 		err := db.QueryRow("insert into movies(movieID, movieName) values($1, $2) returning id;", movieID, movieName).Scan(&lastInsertID)
// 		if err != nil {
// 			response := writeErrResponse(http.StatusBadRequest, "request denied (bad query)", EmptyData{})
// 			log.Fatalln(response)
// 		}

// 		movies = append(movies, Movies{
// 			MovieID:   movieID,
// 			MovieName: movieName})

// 		response = writeResponse(http.StatusOK, "movie added to database", movies)
// 		// response = JsonResponse{Status: http.StatusOK, Data: movies, Message: "movie has been added to database"}
// 		w.Header().Add("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(response)
// 	}
// }

// func DeleteMovie(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	movieID := params["movieid"]

// 	if movieID == "" {
// 		response := writeResponse(http.StatusBadRequest, "no movie id specified", EmptyData{})
// 		w.Header().Add("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(response)
// 	} else {
// 		db := getClientDB()

// 		_, err := db.Exec("delete from movies where movieid = $1", movieID)
// 		if err != nil {
// 			writeErrResponse(http.StatusBadRequest, "request denied - bad query", EmptyData{})
// 		}

// 		fmt.Println("deleted a movie with id: " + movieID)

// 		response := writeResponse(http.StatusOK, "data has been deleted", EmptyData{})
// 		w.Header().Add("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(response)
// 	}
// }

// func GetMovieByID(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	var movies []Movies

// 	movieID := params["movieid"]

// 	if movieID == "" {
// 		response := writeResponse(http.StatusBadRequest, "no movie id specified", EmptyData{})
// 		w.Header().Add("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(response)
// 	} else {
// 		db := getClientDB()

// 		var movieName string

// 		row, err := db.Query("select * from movies where movieid = $1", movieID)
// 		if err != nil {
// 			response := writeErrResponse(http.StatusBadRequest, "request denied (bad query)", EmptyData{})
// 			log.Fatalln(response)
// 		}

// 		for row.Next() {
// 			var id int

// 			err = row.Scan(&id, &movieID, &movieName)
// 			if err != nil {
// 				response := writeErrResponse(http.StatusNotFound, "no data in database", EmptyData{})
// 				log.Fatalln(response)
// 			}
// 		}

// 		movies = append(movies, Movies{
// 			MovieID:   movieID,
// 			MovieName: movieName,
// 		})

// 		response := writeResponse(http.StatusOK, "data successfully fetched", movies)

// 		fmt.Println(response)
// 		w.Header().Add("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(response)
// 	}
// }
