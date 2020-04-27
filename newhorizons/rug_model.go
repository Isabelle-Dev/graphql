package newhorizons

// Rug represents  a rug object that graphql is able to parse and display
type Rug struct {
	Name        string
	Image       string
	DIY         string
	Buy         int
	Sell        int
	Color       []string
	Source      string
	SourceNotes string
	Concept     []string
	HHASeries   string
	Tag         string
	Catalog     string
}

// RugEntry represents a rug entry in the postgres database
type RugEntry struct {
	Name        string
	Image       string
	DIY         string
	Buy         int
	Sell        int
	Color1      string `gorm:"column:color1"`
	Color2      string `gorm:"column:color2"`
	Source      string
	SourceNotes string `gorm:"column:sourcenotes"`
	HHAConcept1 string `gorm:"column:hhaconcept1"`
	HHAConcept2 string `gorm:"column:hhaconcept2"`
	HHASeries   string `gorm:"column:hhaseries"`
	Tag         string
	Catalog     string
}

// ToGraphQL (RugEntry) formats the entry into an object that graphql will be
// able to parse
func (r RugEntry) ToGraphQL(color, concept []string) *Rug {
	return &Rug{
		Name:        r.Name,
		Image:       r.Image,
		DIY:         r.DIY,
		Buy:         r.Buy,
		Sell:        r.Sell,
		Color:       color,
		Source:      r.Source,
		SourceNotes: r.SourceNotes,
		Concept:     concept,
		HHASeries:   r.HHASeries,
		Tag:         r.Tag,
		Catalog:     r.Catalog,
	}
}
