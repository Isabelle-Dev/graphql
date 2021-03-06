package bugs

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB, month string) []*newhorizons.Bug {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var bugs []newhorizons.BugEntry
	db.Raw(dbStr).Find(&bugs)
	if len(bugs) == 0 {
		return nil
	}
	return toBugSlice(bugs, month)
}
