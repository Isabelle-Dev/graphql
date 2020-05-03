package newhorizons

// Item represents an item entry that will be displayed via graphql
type Item struct {
	Name             string
	Variants         []Variant
	BodyTitle        string
	PatternTitle     string
	BodyCustomize    string
	KitCost          int
	Sell             int
	Interact         string
	Source           string
	Tag              string
	HHAConcepts      []string
	HHASeries        string
	HHASet           string
	DIY              string
	PatternCustomize string
	Buy              int
}

// Variant represents variation differences according to each
// item entry
type Variant struct {
	Image     string
	Pattern   string
	Colors    []string
	Variation string
}

// ItemEntry represents an item from the housewares table in postgres
type ItemEntry struct {
	Name             string `gorm:"not null"`
	Image            string `gorm:"not null"`
	Variation        string `gorm:"not null"`
	BodyTitle        string `gorm:"column:bodytitle"`
	Pattern          string
	PatternTitle     string `gorm:"column:patterntitle"`
	DIY              string
	BodyCustomize    string `gorm:"column:bodycustomize"`
	PatternCustomize string `gorm:"column:patterncustomize"`
	KitCost          int    `gorm:"column:kitcost"`
	Buy              int
	Sell             int
	Color1           string `gorm:"column:color1"`
	Color2           string `gorm:"column:color2"`
	Source           string
	SourceNotes      string `gorm:"column:sourcenotes"`
	HHAConcept1      string `gorm:"column:hhaconcept1"`
	HHAConcept2      string `gorm:"column:hhaconcept2"`
	HHASeries        string `gorm:"column:hhaseries"`
	HHASet           string `gorm:"column:hhaset"`
	Interact         string
	Tag              string
	Catalog          string
}

// ToGraphQL (ItemEntry) formats the entry into an object that graphql will be
// able to parse
func (ie ItemEntry) ToGraphQL(v []Variant, h []string) *Item {
	return &Item{
		Name:             ie.Name,
		Variants:         v,
		BodyTitle:        ie.BodyTitle,
		PatternTitle:     ie.PatternTitle,
		BodyCustomize:    ie.BodyCustomize,
		KitCost:          ie.KitCost,
		Sell:             ie.Sell,
		Interact:         ie.Interact,
		Source:           ie.Source,
		Tag:              ie.Tag,
		HHAConcepts:      h,
		HHASeries:        ie.HHASeries,
		HHASet:           ie.HHASet,
		DIY:              ie.DIY,
		PatternCustomize: ie.PatternCustomize,
		Buy:              ie.Buy,
	}
}
