package resolvers

type PersonResolver struct {
	p *Person
}

func (r *PersonResolver) FirstName() *string {
	return &r.p.FirstName
}

func (r *PersonResolver) LastName() *string {
	return &r.p.LastName
}
