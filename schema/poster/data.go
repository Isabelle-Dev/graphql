package poster

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB) []*newhorizons.Poster {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var posters []newhorizons.PosterEntry
	db.Raw(dbStr).Find(&posters)
	if len(posters) == 0 {
		return nil
	}
	return toPosterSlice(posters)
}
