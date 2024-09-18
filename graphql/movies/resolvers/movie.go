package resolvers

type MovieResolver struct {
	m *Movie
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
