package newhorizons

// Photo represents a photo entry graphql can parse and display
type Photo struct {
	Name         string
	Variants     []Variant
	BodyTitle    string
	PatternTitle string
	DIY          string
	Customize    string
	KitCost      int
	Buy          int
	Sell         int
	Source       string
}

// PhotoEntry represents a photo entry in the postgres database
type PhotoEntry struct {
	Name         string
	Image        string
	Variation    string
	BodyTitle    string `gorm:"column:bodytitle"`
	Pattern      string
	PatternTitle string `gorm:"column:patterntitle"`
	DIY          string
	Customize    string
	KitCost      int `gorm:"column:kitcost"`
	Buy          int
	Sell         int
	Color1       string `gorm:"column:color1"`
	Color2       string `gorm:"column:color2"`
	Source       string
}

// ToGraphQL (PhotoEntry) turns a clothes entry from database to a Clothes object
// graphql is able to parse
func (pe PhotoEntry) ToGraphQL(v []Variant) *Photo {
	return &Photo{
		Name:         pe.Name,
		Variants:     v,
		BodyTitle:    pe.BodyTitle,
		PatternTitle: pe.PatternTitle,
		DIY:          pe.DIY,
		Customize:    pe.Customize,
		KitCost:      pe.KitCost,
		Buy:          pe.Buy,
		Sell:         pe.Sell,
		Source:       pe.Source,
	}
}
