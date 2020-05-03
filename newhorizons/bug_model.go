package newhorizons

// Bug represents a bug object graphql can parse and display
type Bug struct {
	Name      string
	Image     string
	Sell      int
	Location  string
	Weather   string
	Rarity    string
	StartTime string
	EndTime   string
	Color     []string
}

// BugEntry represents a bug object in the database
type BugEntry struct {
	Name      string
	Image     string
	Sell      int
	Location  string
	Weather   string
	Rarity    string
	StartTime string `gorm:"column:starttime"`
	EndTime   string `gorm:"column:endtime"`
	Color1    string `gorm:"color1"`
	Color2    string `gorm:"color2"`
}

// ToGraphQL (BugEntry) converts a BugEntry object into a Bug object
func (be BugEntry) ToGraphQL(color []string) *Bug {
	return &Bug{
		Name:      be.Name,
		Image:     be.Image,
		Sell:      be.Sell,
		Location:  be.Location,
		Weather:   be.Weather,
		Rarity:    be.Rarity,
		StartTime: be.StartTime,
		EndTime:   be.EndTime,
		Color:     color,
	}
}
