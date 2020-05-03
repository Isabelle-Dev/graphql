package fossil

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func execute(dbStr string, db *gorm.DB) []*newhorizons.Fossil {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var fossil []newhorizons.FossilEntry
	db.Raw(dbStr).Find(&fossil)
	if len(fossil) == 0 {
		return nil
	}
	return toFossilSlice(fossil)
}
