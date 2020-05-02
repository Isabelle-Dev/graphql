package newhorizons

// Other represents an other object in the database
// and as an object graphql can understand
type Other struct {
	Name        string
	Image       string
	DIY         string
	Buy         int
	Sell        int
	Source      string
	SourceNotes string `gorm:"column:sourcenotes"`
	Tag         string
}
