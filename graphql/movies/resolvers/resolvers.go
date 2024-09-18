package resolvers

type RootResolver struct {
	db *DB
}

func NewRootResolver(db *DB) *RootResolver {
	return &RootResolver{db: db}
}

func (r *RootResolver) Hello() string { return "Hello, world!!!" }

func (r *RootResolver) Movies() *[]*MovieResolver {
	resolvers := make([]*MovieResolver, len(r.db.Records))
	for i, mr := range r.db.Records {
		resolvers[i] = &MovieResolver{m: mr.ToMovie()}
	}
	return &resolvers
}

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

func (r *RootResolver) CreateMovie(args struct{ Input *MovieInput }) *MovieResolver {
	mr, err := r.db.CreateMovie(args.Input)
	if err != nil {
		return &MovieResolver{nil}
	}
	return &MovieResolver{m: mr.ToMovie()}
}
