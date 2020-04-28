package villager

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func execute(dbStr string, db *gorm.DB) []*newhorizons.Villager {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var villager []newhorizons.VillagerEntry
	db.Raw(dbStr).Find(&villager)
	if len(villager) == 0 {
		return nil
	}
	return toVillagerSlice(villager)
}
