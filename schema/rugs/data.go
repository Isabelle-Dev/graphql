package rug

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB) []*newhorizons.Rug {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var rugs []newhorizons.RugEntry
	db.Raw(dbStr).Find(&rugs)
	if len(rugs) == 0 {
		return nil
	}
	return toRugSlice(rugs)
}
