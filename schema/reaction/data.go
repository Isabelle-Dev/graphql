package reaction

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB) []newhorizons.ReactionEntry {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var r []newhorizons.ReactionEntry
	db.Raw(dbStr).Find(&r)
	if len(r) == 0 {
		return nil
	}
	return r
}
