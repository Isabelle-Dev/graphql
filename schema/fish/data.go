package fish

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB, month string) []*newhorizons.Fish {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var fish []newhorizons.FishEntry
	db.Raw(dbStr).Find(&fish)
	if len(fish) == 0 {
		return nil
	}
	return toFishSlice(fish, month)
}
