package newhorizons

// Art represents an art object graphql can parse and display
type Art struct {
	Name     string
	Type     []ArtType
	Buy      int
	Source   string
	Concept  []string
	Tag      string
	Category string
}

// ArtType represents an artwork in terms of genuine vs fake ones
type ArtType struct {
	Image   string
	Sell    int
	Genuine string
}

// ArtEntry represents an art entry in the database
type ArtEntry struct {
	Name        string
	Image       string
	Genuine     string
	Category    string
	Buy         int
	Sell        int
	Source      string
	HHAConcept1 string `gorm:"column:hhaconcept1"`
	HHAConcept2 string `gorm:"column:hhaconcept2"`
	Tag         string
}

// ToGraphQL (ArtEntry) turns a art entry from database to a Art object
// graphql is able to parse
func (ae ArtEntry) ToGraphQL(t []ArtType, c []string) *Art {
	return &Art{
		Name:     ae.Name,
		Type:     t,
		Buy:      ae.Buy,
		Source:   ae.Source,
		Concept:  c,
		Tag:      ae.Tag,
		Category: ae.Category,
	}
}
