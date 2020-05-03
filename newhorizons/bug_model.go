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
	StartTime string
	EndTime   string
	Color1    string
	Color2    string
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
