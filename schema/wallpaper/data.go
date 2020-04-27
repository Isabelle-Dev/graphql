package wallpaper

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB) []*newhorizons.Wallpaper {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var wallpaper []newhorizons.WallpaperEntry
	db.Raw(dbStr).Find(&wallpaper)
	if len(wallpaper) == 0 {
		return nil
	}
	return toWallpaperSlice(wallpaper, db)
}
