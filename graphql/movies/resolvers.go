package movies

type RootResolver struct {
	db *DB
}

//type MoviesResolver struct {
//	Movies []*MovieResolver
//}

type MovieResolver struct {
	m *Movie
}

type PersonResolver struct {
	p *Person
}

func NewRootResolver(db *DB) *RootResolver {
	return &RootResolver{db: db}
}

func (r *RootResolver) Hello() string { return "Hello, world!!!" }

//func (r *RootResolver) Movies() *MoviesResolver {
//	resolvers := make([]*MovieResolver, len(r.db.records))
//	for i, mr := range r.db.records {
//		resolvers[i] = &MovieResolver{m: mr.ToMovie()}
//	}
//	return &MoviesResolver{Movies: resolvers}
//}

func (r *RootResolver) Movie(args struct{ Id int32 }) *MovieResolver {
	mr, err := r.db.GetMovie(int(args.Id))
	if err != nil {
		return &MovieResolver{nil}
	}
	return &MovieResolver{m: mr.ToMovie()}
}

func (r *RootResolver) MovieFromTitle(args struct{ Title string }) *MovieResolver {
	mr, err := r.db.GetMovieFromTitle(args.Title)
	if err != nil {
		return &MovieResolver{nil}
	}
	return &MovieResolver{m: mr.ToMovie()}
}

func (r *MovieResolver) Id() int32 {
	return rune(r.m.Id)
}

func (r *MovieResolver) Title() string {
	return r.m.Title
}

func (r *MovieResolver) Director() *PersonResolver {
	if r.m.Director == nil {
		return nil
	}
	return &PersonResolver{p: r.m.Director}
}

func (r *MovieResolver) Actors() *[]*PersonResolver {
	if r.m.Actors == nil {
		return nil
	}
	resolvers := make([]*PersonResolver, len(r.m.Actors))
	for i, actor := range r.m.Actors {
		resolvers[i] = &PersonResolver{p: actor}
	}
	return &resolvers
}

func (r *MovieResolver) Plot() *string {
	return &r.m.Plot
}

func (r *MovieResolver) Genres() *[]string {
	return &r.m.Genres
}

func (r *MovieResolver) Year() *int32 {
	year := int32(r.m.Year)
	return &year
}

func (r *MovieResolver) Country() *string {
	return &r.m.Country
}

func (r *MovieResolver) Language() *string {
	return &r.m.Language
}

func (r *MovieResolver) Rating() *float64 {
	return &r.m.Rating
}

func (r *PersonResolver) FirstName() *string {
	return &r.p.FirstName
}

func (r *PersonResolver) LastName() *string {
	return &r.p.LastName
}
