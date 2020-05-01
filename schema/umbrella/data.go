package umbrella

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func execute(dbStr string, db *gorm.DB) []*newhorizons.Umbrella {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var um []newhorizons.UmbrellaEntry
	db.Raw(dbStr).Find(&um)
	if len(um) == 0 {
		return nil
	}
	return toUmbrellaSlice(um)
}
