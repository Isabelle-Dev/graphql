package newhorizons

// Item represents an item entry that will be displayed via graphql
type Item struct {
	Name             string
	Image            []string
	Variation        []string
	BodyTitle        string
	Pattern          []string
	PatternTitle     string
	BodyCustomize    string
	KitCost          string
	Sell             int
	Interact         string
	Source           string
	Tag              string
	HHAConcepts      []string
	HHASeries        string
	HHASet           string
	DIY              string
	PatternCustomize string
	Buy              string
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
	KitCost          string `gorm:"column:kitcost"`
	Buy              string
	Sell             int
	Color1           string `gorm:"column:color1"`
	Color2           string `gorm:"column:color2"`
	Size             string
	Source           string
	SourceNotes      string `gorm:"column:sourcenotes"`
	Version          string
	HHAConcept1      string `gorm:"column:hhaconcept1"`
	HHAConcept2      string `gorm:"column:hhaconcept2"`
	HHASeries        string `gorm:"column:hhaseries"`
	HHASet           string `gorm:"column:hhaset"`
	Interact         string
	Tag              string
	Catalog          string
	Filename         string
	VariantID        string `gorm:"column:variantid"`
	InternalID       int    `gorm:"column:internalid"`
}

// ToGraphQL (ItemEntry) formats the entry into an object that graphql will be
// able to parse
func (ie ItemEntry) ToGraphQL(v, p, i, h []string) *Item {
	return &Item{
		Name:             ie.Name,
		Image:            i,
		Variation:        v,
		BodyTitle:        ie.BodyTitle,
		Pattern:          p,
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
