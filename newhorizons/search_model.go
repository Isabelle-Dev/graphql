package newhorizons

// BugEntry represents a database entry of a bug
// in the postgres database
type BugEntry struct {
	Num                  int    `gorm:"not null;column:num"`
	Name                 string `gorm:"not null;column:name"`
	Image                string `gorm:"not null;column:image"`
	House                string `gorm:"not null;column:house"`
	Sell                 int    `gorm:"not null;column:sell"`
	Where                string `gorm:"not null;column:location"`
	Weather              string `gorm:"not null;column:weather"`
	Rarity               string `gorm:"not null;column:rarity"`
	StartTime            string `gorm:"not null;column:starttime"`
	EndTime              string `gorm:"not null;column:endtime"`
	Jan                  string `gorm:"not null;column:jan"`
	Feb                  string `gorm:"not null;column:feb"`
	Mar                  string `gorm:"not null;column:mar"`
	Apr                  string `gorm:"not null;column:apr"`
	May                  string `gorm:"not null;column:may"`
	Jun                  string `gorm:"not null;column:jun"`
	Jul                  string `gorm:"not null;column:jul"`
	Aug                  string `gorm:"not null;column:aug"`
	Sep                  string `gorm:"not null;column:sep"`
	Oct                  string `gorm:"not null;column:oct"`
	Nov                  string `gorm:"not null;column:nov"`
	Dec                  string `gorm:"not null;column:dec"`
	Color1               string `gorm:"not null;column:color1"`
	Color2               string `gorm:"not null;column:color2"`
	CritterpediaFilename string `gorm:"-"`
	ItemFilename         string `gorm:"not null;column:itemfilename"`
	InternalID           int    `gorm:"not null;column:internalid"`
}

// FishEntry represents all the fields which represent a fish entry
type FishEntry struct {
	Num                  int    `gorm:"not null;column:num"`
	Name                 string `gorm:"not null;column:name"`
	Image                string `gorm:"not null;column:image"`
	House                string `gorm:"not null;column:house"`
	Sell                 int    `gorm:"not null;column:sell"`
	Where                string `gorm:"not null;column:location"`
	Shadow               string `gorm:"not null;column:shadow"`
	Rarity               string `gorm:"not null;column:rarity"`
	RainSnowUp           string `gorm:"not null;column:rainsnowup"`
	StartTime            string `gorm:"not null;column:starttime"`
	EndTime              string `gorm:"not null;column:endtime"`
	Jan                  string `gorm:"not null;column:jan"`
	Feb                  string `gorm:"not null;column:feb"`
	Mar                  string `gorm:"not null;column:mar"`
	Apr                  string `gorm:"not null;column:apr"`
	May                  string `gorm:"not null;column:may"`
	Jun                  string `gorm:"not null;column:jun"`
	Jul                  string `gorm:"not null;column:jul"`
	Aug                  string `gorm:"not null;column:aug"`
	Sep                  string `gorm:"not null;column:sep"`
	Oct                  string `gorm:"not null;column:oct"`
	Nov                  string `gorm:"not null;column:nov"`
	Dec                  string `gorm:"not null;column:dec"`
	Color1               string `gorm:"not null;column:color1"`
	Color2               string `gorm:"not null;column:color2"`
	Size                 string `gorm:"not null;column:size"`
	LightingType         string `gorm:"not null;column:lightingtype"`
	CritterpediaFilename string `gorm:"-"`
	ItemFilename         string `gorm:"not null;column:itemfilename"`
	InternalID           int    `gorm:"not null;column:internalid"`
}
