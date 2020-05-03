package newhorizons

// Fossil represents a fossil entry graphql can parse and display
type Fossil struct {
	Name     string
	Image    string
	Buy      int
	Sell     int
	Color    []string
	Source   string
	Interact string
	Catalog  string
}

// FossilEntry represents a fossil entry in the database
type FossilEntry struct {
	Name     string
	Image    string
	Buy      int
	Sell     int
	Color1   string
	Color2   string
	Source   string
	Interact string
	Catalog  string
}

// ToGraphQL (FossilEntry) turns a FossilEntry into a Fossil object
func (fe FossilEntry) ToGraphQL(color []string) *Fossil {
	return &Fossil{
		Name:     fe.Name,
		Image:    fe.Image,
		Buy:      fe.Buy,
		Sell:     fe.Sell,
		Color:    color,
		Source:   fe.Source,
		Interact: fe.Interact,
		Catalog:  fe.Catalog,
	}
}
