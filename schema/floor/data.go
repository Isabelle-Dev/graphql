package floor

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func execute(dbStr string, db *gorm.DB) []*newhorizons.Floor {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var floors []newhorizons.FloorEntry
	db.Raw(dbStr).Find(&floors)
	if len(floors) == 0 {
		return nil
	}
	return toFloorSlice(floors)
}
