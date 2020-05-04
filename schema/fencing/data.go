package fencing

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB) []newhorizons.FenceEntry {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var fences []newhorizons.FenceEntry
	db.Raw(dbStr).Find(&fences)
	if len(fences) == 0 {
		return nil
	}
	return fences
}
