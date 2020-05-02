package newhorizons

// NookMiles represents a NookMile object in both the database
// and as a graphql entry
type NookMiles struct {
	Name      string
	Image     string
	NookMiles int `gorm:"column:nookmiles"`
	Category  string
}
