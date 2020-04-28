package music

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

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
