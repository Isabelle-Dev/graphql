package music

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB) []*newhorizons.Music {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var music []newhorizons.MusicEntry
	db.Raw(dbStr).Find(&music)
	if len(music) == 0 {
		return nil
	}
	return toMusicSlice(music)
}
