package newhorizons

// Poster represents a poster that graphql can parse and display
type Poster struct {
	Name   string
	Image  string
	Buy    int
	Sell   int
	Color  []string
	Source string
}

// PosterEntry represents a poster in the database
type PosterEntry struct {
	Name   string
	Image  string
	Buy    int
	Sell   int
	Color1 string `gorm:"column:color1"`
	Color2 string `gorm:"column:color2"`
	Source string
}

// ToGraphQL (PosterEntry) turns a poster entry from database to a Poster object
// graphql is able to parse
func (pe PosterEntry) ToGraphQL(c []string) *Poster {
	return &Poster{
		Name:   pe.Name,
		Image:  pe.Image,
		Buy:    pe.Buy,
		Sell:   pe.Sell,
		Color:  c,
		Source: pe.Source,
	}
}
