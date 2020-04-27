package newhorizons

// Floor represents a floor entry graphql is able to parse and display
type Floor struct {
	Name        string
	Image       string
	VFX         string
	DIY         string
	Buy         int
	Sell        int
	Color       []string
	Source      string
	SourceNotes string
	Concepts    []string
	Tag         string
	Catalog     string
}

// FloorEntry represents a floor entry from the database
type FloorEntry struct {
	Name        string
	Image       string
	VFX         string
	DIY         string
	Buy         int
	Sell        int
	Color1      string `gorm:"column:color1"`
	Color2      string `gorm:"column:color2"`
	Source      string
	SourceNotes string `gorm:"column:sourcenotes"`
	HHAConcept1 string `gorm:"column:hhaconcept1"`
	HHAConcept2 string `gorm:"column:hhaconcept2"`
	Tag         string `gorm:"column:tag"`
	Catalog     string
}

// ToGraphQL (FloorEntry) turns a clothes entry from database to a Floor object
// graphql is able to parse
func (f FloorEntry) ToGraphQL(color, concepts []string) *Floor {
	return &Floor{
		Name:        f.Name,
		Image:       f.Image,
		VFX:         f.VFX,
		DIY:         f.DIY,
		Buy:         f.Buy,
		Sell:        f.Sell,
		Color:       color,
		Source:      f.Source,
		SourceNotes: f.SourceNotes,
		Concepts:    concepts,
		Tag:         f.Tag,
		Catalog:     f.Catalog,
	}
}
