package nookmiles

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func execute(dbStr string, db *gorm.DB) []newhorizons.NookMiles {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var nm []newhorizons.NookMiles
	db.Raw(dbStr).Find(&nm)
	if len(nm) == 0 {
		return nil
	}
	return nm
}
