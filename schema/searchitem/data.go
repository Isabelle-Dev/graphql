package searchitem

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// findAllDoorHanging retrieves all entries that are counted as a hanging item on doors
func findAllDoorHanging(tablename string, db *gorm.DB) []*newhorizons.Item {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var items []newhorizons.ItemEntry
	db.Table(tablename).Where("tag LIKE ?", "%Door Decor%").Find(&items)
	if len(items) == 0 {
		return nil
	}
	return toItemSlice(items, db)
}

// findByConcept retrieves an extry from the item table by a given concept filter
func findByConcept(concept, tablename string, db *gorm.DB) []*newhorizons.Item {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var items []newhorizons.ItemEntry
	db.Table(tablename).Where("hhaconcept1 = ? OR hhaconcept2 = ?", concept, concept).Find(&items)
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
	v, p, i, h := extractVPIH(items)
	return items[0].ToGraphQL(v, p, i, h)
}

// findByColor retrieves a list of entries from the item table by a given color
// variation - this option excludes pattern variations
func findByColor(color, tablename string, db *gorm.DB) []*newhorizons.Item {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var items []newhorizons.ItemEntry
	db.Table(tablename).Where("variation = ? AND pattern = ?", color, "NA").Find(&items)
	if len(items) == 0 {
		return nil
	}
	return toItemSlice(items, db)
}
