package item

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB) []*newhorizons.Item {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var items []newhorizons.ItemEntry
	db.Raw(dbStr).Find(&items)
	if len(items) == 0 {
		return nil
	}
	return toItemSlice(items, db)
}

// findByName retrieves an entry from the item table by
// a given name
func findByName(name, tablename string, db *gorm.DB) *newhorizons.Item {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var items []newhorizons.ItemEntry
	db.Table(tablename).Where("name = ?", name).Find(&items)
	if len(items) == 0 {
		return nil
	}
	v, h := extractVPIH(items)
	return items[0].ToGraphQL(v, h)
}
