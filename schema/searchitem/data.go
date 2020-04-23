package searchitem

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// findByName retrieves an entry from the item table by
// a given name
func findByName(name, tablename string, db *gorm.DB) *newhorizons.Item {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var items []newhorizons.ItemEntry
	err := db.Table(tablename).Where("name = ?", name).Find(&items).Error
	if err != nil {
		return nil
	}
	v, p, i, h := extractVPIH(items)
	return items[0].ToGraphQL(v, p, i, h)
}

// findByColor retrieves a list of entries from the item table by a given color
// variation - this option excludes pattern variations
func findByColor(color, tablename string, db *gorm.DB) []*newhorizons.Item {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var items []newhorizons.ItemEntry
	err := db.Table(tablename).Where("variation = ? AND pattern = ?", color, "NA").Find(&items).Error
	if err != nil {
		return nil
	}
	var ret []*newhorizons.Item
	for _, i := range items {
		n := findByName(i.Name, "housewares", db)
		ret = append(ret, i.ToGraphQL(n.Variation, n.Pattern, n.Image, n.HHAConcepts))
	}
	return ret
}
