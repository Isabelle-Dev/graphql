package fish

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func execute(dbStr string, db *gorm.DB) []*newhorizons.Fish {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var fish []newhorizons.FishEntry
	db.Raw(dbStr).Find(&fish)
	if len(fish) == 0 {
		return nil
	}
	return toFishSlice(fish)
}
