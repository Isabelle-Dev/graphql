package newhorizons

// Fish represents a fish object graphql can parse and display
type Fish struct {
	Name      string
	Image     string
	Sell      int
	Location  string
	Shadow    string
	Rarity    string
	CatchUp   string
	StartTime string
	EndTime   string
	Color     []string
}

// FishEntry represents a fish object in the database
type FishEntry struct {
	Name      string
	Image     string
	Sell      int
	Location  string
	Shadow    string
	Rarity    string
	CatchUp   string
	StartTime string
	EndTime   string
	Color1    string
	Color2    string
}

// ToGraphQL (FishEntry) converts a FishEntry object into a Fish object
func (fe FishEntry) ToGraphQL(color []string) *Fish {
	return &Fish{
		Name:      fe.Name,
		Image:     fe.Image,
		Sell:      fe.Sell,
		Location:  fe.Location,
		Shadow:    fe.Shadow,
		Rarity:    fe.Rarity,
		CatchUp:   fe.CatchUp,
		StartTime: fe.StartTime,
		EndTime:   fe.EndTime,
		Color:     color,
	}
}
