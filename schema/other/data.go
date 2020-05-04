package other

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB) []newhorizons.Other {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var other []newhorizons.Other
	db.Raw(dbStr).Find(&other)
	if len(other) == 0 {
		return nil
	}
	return other
}
