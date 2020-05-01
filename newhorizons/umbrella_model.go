package newhorizons

// Umbrella represents an umbrella object graphql can parse and display
type Umbrella struct {
	Name        string
	Image       string
	DIY         string
	Buy         int
	Sell        int
	Color       []string
	Source      string
	SourceNotes string
	Catalog     string
}

// UmbrellaEntry represents an umbrella in the database
type UmbrellaEntry struct {
	Name        string
	Image       string
	DIY         string
	Buy         int
	Sell        int
	Color1      string `gorm:"column:color1"`
	Color2      string `gorm:"column:color2"`
	Source      string
	SourceNotes string `gorm:"column:sourcenotes"`
	Catalog     string
}

// ToGraphQL (UmbrellaEntry) turns an umbrella entry from database to an Umbrella object
// graphql is able to parse
func (ue UmbrellaEntry) ToGraphQL(c []string) *Umbrella {
	return &Umbrella{
		Name:        ue.Name,
		Image:       ue.Image,
		DIY:         ue.DIY,
		Buy:         ue.Buy,
		Sell:        ue.Sell,
		Color:       c,
		Source:      ue.Source,
		SourceNotes: ue.SourceNotes,
		Catalog:     ue.Catalog,
	}
}
