package search

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// findByName retrieves a database entry by a given item name
//
// If none is found, it will return an error
func findByName(item, tablename string, db *gorm.DB) *newhorizons.BugEntry {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var entry newhorizons.BugEntry
	err := db.Table(tablename).Where("name = ?", item).First(&entry).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return &entry
}
