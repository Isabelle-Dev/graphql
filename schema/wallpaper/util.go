package wallpaper

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// buildSchema retrieves nested data from a wallpaper entry and formats them into
// an object graphql will use
func buildSchema(w newhorizons.WallpaperEntry) (newhorizons.VFXT, newhorizons.Windows, newhorizons.Curtains, []string, []string) {
	vfx := newhorizons.VFXT{
		VFX:     w.VFX,
		VFXType: w.VFXType,
	}
	win := newhorizons.Windows{
		WindowType:  w.WindowType,
		WindowColor: w.WindowColor,
	}
	cur := newhorizons.Curtains{
		CurtainType:  w.CurtainType,
		CurtainColor: w.CurtainColor,
	}
	var color []string
	color = append(color, w.Color1)
	if !exists(w.Color2, color) {
		color = append(color, w.Color2)
	}
	var concepts []string
	concepts = append(concepts, w.HHAConcept1)
	if !exists(w.Color2, concepts) {
		concepts = append(concepts, w.HHAConcept2)
	}
	return vfx, win, cur, color, concepts
}

// turns WallpaperEntry into Wallpaper
func toWallpaperSlice(w []newhorizons.WallpaperEntry, db *gorm.DB) []*newhorizons.Wallpaper {
	var ret []*newhorizons.Wallpaper
	for _, i := range w {
		v, w, cur, color, concepts := buildSchema(i)
		ret = append(ret, i.ToGraphQL(v, w, cur, color, concepts))
	}
	return ret
}

func exists(c string, con []string) bool {
	for _, i := range con {
		if c == i {
			return true
		}
	}
	return false
}