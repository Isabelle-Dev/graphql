package newhorizons

// Music represents a music object that graphql is able to parse and display
type Music struct {
	Name        string
	Image       string
	Buy         int
	Sell        int
	Color       []string
	Source      string
	SourceNotes string
}

// MusicEntry represents a music object from the database
type MusicEntry struct {
	Name        string
	Image       string
	Buy         int
	Sell        int
	Color1      string `gorm:"column:color1"`
	Color2      string `gorm:"column:color2"`
	Source      string
	SourceNotes string `gorm:"column:sourcenotes"`
}

// ToGraphQL (MusicEntry) turns a MusicEntry object into a Music Object
func (me MusicEntry) ToGraphQL(c []string) *Music {
	return &Music{
		Name:        me.Name,
		Image:       me.Image,
		Buy:         me.Buy,
		Sell:        me.Sell,
		Color:       c,
		Source:      me.Source,
		SourceNotes: me.SourceNotes,
	}
}
