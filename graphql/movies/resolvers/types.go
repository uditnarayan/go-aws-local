package resolvers

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

type MovieInput struct {
	Title    string  `json:"title"`
	Country  *string `json:"country"`
	Language *string `json:"language"`
}
