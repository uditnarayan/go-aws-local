package movies

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Movie struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Year     int       `json:"year"`
	Rating   float64   `json:"rating"`
	Director *Person   `json:"director"`
	Actors   []*Person `json:"actors"`
	Genres   []string  `json:"genres"`
	Plot     string    `json:"plot"`
	Country  string    `json:"country"`
	Language string    `json:"language"`
}

type MovieRecord struct {
	Id       int      `json:"id"`
	Title    string   `json:"title"`
	Year     int      `json:"year"`
	Rating   float64  `json:"rating"`
	Director string   `json:"director"`
	Actors   []string `json:"actors"`
	Genres   []string `json:"genres"`
	Plot     string   `json:"plot"`
	Country  string   `json:"country"`
	Language string   `json:"language"`
}

type DB struct {
	records []*MovieRecord
}

func Connect() (*DB, error) {
	res, err := http.Get("https://freetestapi.com/api/v1/movies")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch movies. %v", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read movies. %v", err)
	}
	var movies []*MovieRecord
	err = json.Unmarshal(body, &movies)
	if err != nil {
		return nil, fmt.Errorf("failed to pasre json. %v", err)
	}
	return &DB{records: movies}, nil
}

func (db *DB) GetMovie(id int) (*MovieRecord, error) {
	for _, movie := range db.records {
		if movie.Id == id {
			return movie, nil
		}
	}
	return nil, fmt.Errorf("failed to find movie with id %s", id)
}

func (db *DB) GetMovieFromTitle(title string) (*MovieRecord, error) {
	for _, movie := range db.records {
		if strings.Compare(movie.Title, title) == 0 {
			return movie, nil
		}
	}
	return nil, fmt.Errorf("failed to find movie with title %s", title)
}

func (mr *MovieRecord) ToMovie() *Movie {
	m := &Movie{
		Id:       mr.Id,
		Title:    mr.Title,
		Year:     mr.Year,
		Rating:   mr.Rating,
		Genres:   mr.Genres,
		Plot:     mr.Plot,
		Country:  mr.Country,
		Language: mr.Language,
	}
	dir := strings.Split(mr.Director, " ")
	m.Director = &Person{FirstName: dir[0], LastName: dir[1]}
	m.Actors = make([]*Person, len(mr.Actors))
	for i, actor := range mr.Actors {
		splits := strings.Split(actor, " ")
		m.Actors[i] = &Person{FirstName: splits[0], LastName: splits[1]}
	}
	return m
}
