package newhorizons

// Villager represents a villager entry graphql is able to parse and
// display
type Villager struct {
	Name        string
	Image       string
	Species     string
	Gender      string
	Personality string
	Birthday    string
	Catchphrase string
	Style       []string
	Color       []string
}

// VillagerEntry represents a villager in the database
type VillagerEntry struct {
	Name        string
	Image       string
	Species     string
	Gender      string
	Personality string
	Birthday    string
	Catchphrase string
	Style1      string `gorm:"column:style1"`
	Style2      string `gorm:"column:style2"`
	Color1      string `gorm:"column:color1"`
	Color2      string `gorm:"column:color2"`
}

// ToGraphQL (VillagerEntry) turns a villager entry from database to a Villager object
// graphql is able to parse
func (ve VillagerEntry) ToGraphQL(s, c []string) *Villager {
	return &Villager{
		Name:        ve.Name,
		Image:       ve.Image,
		Species:     ve.Species,
		Gender:      ve.Gender,
		Personality: ve.Personality,
		Birthday:    ve.Birthday,
		Catchphrase: ve.Catchphrase,
		Style:       s,
		Color:       c,
	}
}
