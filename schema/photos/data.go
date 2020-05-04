package photos

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB) []*newhorizons.Photo {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var photos []newhorizons.PhotoEntry
	db.Raw(dbStr).Find(&photos)
	if len(photos) == 0 {
		return nil
	}
	return toPhotoSlice(photos, db)
}

// findByName retrieves an entry from the photo table by
// a given name
func findByName(name, tablename string, db *gorm.DB) *newhorizons.Photo {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var photos []newhorizons.PhotoEntry
	db.Table(tablename).Where("name = ?", name).Find(&photos)
	if len(photos) == 0 {
		return nil
	}
	v := extractV(photos)
	return photos[0].ToGraphQL(v)
}
