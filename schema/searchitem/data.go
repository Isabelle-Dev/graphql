package searchitem

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// findByName retrieves a list of entries from the item table which contains
// a given name
func findByName(name, tablename string, db *gorm.DB) *[]newhorizons.Item {
	var items []newhorizons.Item
	err := db.Table(tablename).Where("name = ?", name).Find(&items).Error
	if err != nil {
		return nil
	}
	return &items
}
