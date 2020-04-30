package newhorizons

// Clothes represents a Clothes item that will be displayed via graphql
type Clothes struct {
	Name        string
	Variants    []ClothesVariant
	DIY         string
	Buy         int
	Sell        int
	Source      string
	SourceNotes string
	Style       string
	Themes      []string
	Catalog     string
}

// ClothesVariant represents variation data for clothes entries
type ClothesVariant struct {
	Image     string
	Variation string
	Color     []string
}

// ClothesEntry represents data we get from SQL database
type ClothesEntry struct {
	Name        string
	Image       string
	Variation   string
	DIY         string
	Buy         int
	Sell        int
	Color1      string `gorm:"column:color1"`
	Color2      string `gorm:"column:color2"`
	Source      string
	SourceNotes string `gorm:"column:sourcenotes"`
	Style       string
	Themes      string `gorm:"column:labelthemes"`
	Catalog     string
}

// ToGraphQL (ClothesEntry) turns a clothes entry from database to a Clothes object
// graphql is able to parse
func (c ClothesEntry) ToGraphQL(v []ClothesVariant, t []string) *Clothes {
	return &Clothes{
		Name:        c.Name,
		Variants:    v,
		DIY:         c.DIY,
		Buy:         c.Buy,
		Sell:        c.Sell,
		Source:      c.Source,
		SourceNotes: c.SourceNotes,
		Style:       c.Style,
		Themes:      t,
		Catalog:     c.Catalog,
	}
}
