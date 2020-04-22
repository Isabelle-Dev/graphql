package searchitem

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// findByName retrieves an entry from the item table by
// a given name
func findByName(name, tablename string, db *gorm.DB) *newhorizons.ModifiedItem {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var items []newhorizons.Item
	err := db.Table(tablename).Where("name = ?", name).Find(&items).Error
	if err != nil {
		return nil
	}
	v, p, i := extractVPI(items)
	return toModifiedItem(items[0], v, p, i)
}

// findByColor retrieves a list of entries from the item table by a given color
// variation - this option excludes pattern variations
func findByColor(color, tablename string, db *gorm.DB) *[]newhorizons.Item {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var items []newhorizons.Item
	err := db.Table(tablename).Where("variation = ? AND pattern = ?", color, "NA").Find(&items).Error
	if err != nil {
		return nil
	}
	return &items
}
