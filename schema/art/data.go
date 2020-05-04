package art

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB) []*newhorizons.Art {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var art []newhorizons.ArtEntry
	db.Raw(dbStr).Find(&art)
	if len(art) == 0 {
		return nil
	}
	return toArtSlice(art, db)
}

// Utility database func mainly used to return a singular object from
// multiple variants
func findByName(name, tablename string, db *gorm.DB) *newhorizons.Art {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var art []newhorizons.ArtEntry
	db.Table(tablename).Where("name = ?", name).Find(&art)
	if len(art) == 0 {
		return nil
	}
	v := extractType(art)
	return art[0].ToGraphQL(v)
}
