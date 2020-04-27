package wallpaper

import "github.com/Isabelle-Dev/isabelle-graphql/newhorizons"

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

func exists(c string, con []string) bool {
	for _, i := range con {
		if c == i {
			return true
		}
	}
	return false
}
