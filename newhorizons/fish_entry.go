package newhorizons

// Fish represents a fish object graphql can parse and display
type Fish struct {
	Name         string
	Image        string
	Sell         int
	Location     string
	Shadow       string
	TotalCatches int
	CatchUp      string
	MonthTime    []Month
	Color        []string
}

// FishEntry represents a fish object in the database
type FishEntry struct {
	Name         string
	Image        string
	Sell         int
	Location     string
	Shadow       string
	TotalCatches int    `gorm:"column:totalcatches"`
	CatchUp      string `gorm:"column:catchup"`
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

// ToGraphQL (FishEntry) converts a FishEntry object into a Fish object
func (fe FishEntry) ToGraphQL(color []string, month []Month) *Fish {
	return &Fish{
		Name:         fe.Name,
		Image:        fe.Image,
		Sell:         fe.Sell,
		Location:     fe.Location,
		Shadow:       fe.Shadow,
		TotalCatches: fe.TotalCatches,
		CatchUp:      fe.CatchUp,
		MonthTime:    month,
		Color:        color,
	}
}
