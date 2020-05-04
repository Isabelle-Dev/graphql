package construction

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB) []newhorizons.ConstructionEntry {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var c []newhorizons.ConstructionEntry
	db.Raw(dbStr).Find(&c)
	if len(c) == 0 {
		return nil
	}
	return c
}
