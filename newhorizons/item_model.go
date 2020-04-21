package newhorizons

// Item represents an item from the housewares catalog
type Item struct {
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
	SpeakerType      string `gorm:"column:speakertype"`
	LightingType     string `gorm:"column:lightingtype"`
	Catalog          string
	Filename         string
	VariantID        string `gorm:"column:variantid"`
	InternalID       int    `gorm:"column:internalid"`
}
