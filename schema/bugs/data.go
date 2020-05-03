package bugs

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func execute(dbStr string, db *gorm.DB) []*newhorizons.Bug {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var bugs []newhorizons.BugEntry
	db.Raw(dbStr).Find(&bugs)
	if len(bugs) == 0 {
		return nil
	}
	return toBugSlice(bugs)
}
