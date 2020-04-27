package newhorizons

// Wallpaper represents a Wallpaper object graphql can parse
type Wallpaper struct {
	Name        string
	Image       string
	VFXInfo     VFXT
	DIY         string
	Sell        int
	Buy         int
	Source      string
	SourceNotes string
	Catalog     string
	WindowInfo  Windows
	CurtainInfo Curtains
	CeilingType string
	Color       []string
	Concepts    []string
	HHASeries   string
	Tag         string
}

// VFXT represents VFX info
type VFXT struct {
	VFX     string
	VFXType string
}

// Windows represent windows info
type Windows struct {
	WindowType  string
	WindowColor string
}

// Curtains represents curtain info
type Curtains struct {
	CurtainType  string
	CurtainColor string
}

// WallpaperEntry represents a wallpaper object in the postgres database
type WallpaperEntry struct {
	Name         string
	Image        string
	VFX          string `gorm:"column:vfx"`
	VFXType      string `gorm:"column:vfxtype"`
	DIY          string `gorm:"column:diy"`
	Buy          int
	Sell         int
	Source       string
	SourceNotes  string `gorm:"column:sourcenotes"`
	Catalog      string
	WindowType   string `gorm:"column:windowtype"`
	WindowColor  string `gorm:"column:windowcolor"`
	PaneType     string `gorm:"column:panetype"`
	CurtainType  string `gorm:"column:curtaintype"`
	CurtainColor string `gorm:"column:curtaincolor"`
	CeilingType  string `gorm:"column:ceilingtype"`
	Color1       string `gorm:"column:color1"`
	Color2       string `gorm:"column:color2"`
	HHAConcept1  string `gorm:"column:hhaconcept1"`
	HHAConcept2  string `gorm:"column:hhaconcept2"`
	HHASeries    string `gorm:"column:hhaseries"`
	Tag          string
}

// ToGraphQL (WallpaperEntry) turns a clothes entry from database to a Clothes object
// graphql is able to parse
func (w WallpaperEntry) ToGraphQL(v VFXT, win Windows, c Curtains, colors, concepts []string) *Wallpaper {
	return &Wallpaper{
		Name:        w.Name,
		Image:       w.Image,
		VFXInfo:     v,
		DIY:         w.DIY,
		Sell:        w.Sell,
		Buy:         w.Buy,
		Source:      w.Source,
		SourceNotes: w.SourceNotes,
		Catalog:     w.Catalog,
		WindowInfo:  win,
		CurtainInfo: c,
		CeilingType: w.CeilingType,
		Color:       colors,
		Concepts:    concepts,
		HHASeries:   w.HHASeries,
		Tag:         w.Tag,
	}
}
