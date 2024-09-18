package resolvers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type DB struct {
	Records []*MovieRecord
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
	return &DB{Records: movies}, nil
}

func (db *DB) GetMovie(id int) (*MovieRecord, error) {
	for _, movie := range db.Records {
		if movie.Id == id {
			return movie, nil
		}
	}
	return nil, fmt.Errorf("failed to find movie with id %s", id)
}

func (db *DB) GetMovieFromTitle(title string) (*MovieRecord, error) {
	for _, movie := range db.Records {
		if strings.Compare(movie.Title, title) == 0 {
			return movie, nil
		}
	}
	return nil, fmt.Errorf("failed to find movie with title %s", title)
}

func (db *DB) CreateMovie(input *MovieInput) (*MovieRecord, error) {
	lastInsertedId := db.Records[len(db.Records)-1].Id
	mr := &MovieRecord{
		Id:       lastInsertedId + 1,
		Title:    input.Title,
		Country:  *input.Country,
		Language: *input.Language,
	}
	db.Records = append(db.Records, mr)
	return mr, nil
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
	if len(mr.Director) != 0 {
		dir := strings.Split(mr.Director, " ")
		m.Director = &Person{FirstName: dir[0], LastName: dir[1]}
	}
	if len(m.Actors) != 0 {
		m.Actors = make([]*Person, len(mr.Actors))
		for i, actor := range mr.Actors {
			splits := strings.Split(actor, " ")
			m.Actors[i] = &Person{FirstName: splits[0], LastName: splits[1]}
		}
	}
	return m
}
