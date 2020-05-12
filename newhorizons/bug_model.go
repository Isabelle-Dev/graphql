package newhorizons

// Bug represents a bug object graphql can parse and display
type Bug struct {
	Name         string
	Image        string
	Sell         int
	Location     string
	Weather      string
	TotalCatches int
	MonthTime    []Month
	Color        []string
}

// Month represents a specific month and its associated time
type Month struct {
	// Month - separated to NH or SH and month name
	Month string

	// Associated time with month value
	Time string
}

// BugEntry represents a bug object in the database
type BugEntry struct {
	Name         string
	Image        string
	Sell         int
	Location     string
	Weather      string
	TotalCatches int    `gorm:"column:totalcatches"`
	NHJan        string `gorm:"column:nhjan"`
	NHFeb        string `gorm:"column:nhfeb"`
	NHMar        string `gorm:"column:nhmar"`
	NHApr        string `gorm:"column:nhapr"`
	NHMay        string `gorm:"column:nhmay"`
	NHJun        string `gorm:"column:nhjun"`
	NHJul        string `gorm:"column:nhjul"`
	NHAug        string `gorm:"column:nhaug"`
	NHSep        string `gorm:"column:nhsep"`
	NHOct        string `gorm:"column:nhoct"`
	NHNov        string `gorm:"column:nhnov"`
	NHDec        string `gorm:"column:nhdec"`
	SHJan        string `gorm:"column:shjan"`
	SHFeb        string `gorm:"column:shfeb"`
	SHMar        string `gorm:"column:shmar"`
	SHApr        string `gorm:"column:shapr"`
	SHMay        string `gorm:"column:shmay"`
	SHJun        string `gorm:"column:shjun"`
	SHJul        string `gorm:"column:shjul"`
	SHAug        string `gorm:"column:shaug"`
	SHSep        string `gorm:"column:shsep"`
	SHOct        string `gorm:"column:shoct"`
	SHNov        string `gorm:"column:shnov"`
	SHDec        string `gorm:"column:shdec"`
	Color1       string `gorm:"column:color1"`
	Color2       string `gorm:"column:color2"`
}

// ToGraphQL (BugEntry) converts a BugEntry object into a Bug object
func (be BugEntry) ToGraphQL(color []string, month []Month) *Bug {
	return &Bug{
		Name:         be.Name,
		Image:        be.Image,
		Sell:         be.Sell,
		Location:     be.Location,
		Weather:      be.Weather,
		TotalCatches: be.TotalCatches,
		MonthTime:    month,
		Color:        color,
	}
}
