package searchbugandfish

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// findByPrice returns the names of all bug and fish listings that match a given price
func findByPrice(price int, bugTable, fishTable string, db *gorm.DB) *newhorizons.Combined {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var b []newhorizons.BugEntry
	var f []newhorizons.FishEntry
	err := db.Table(bugTable).Where("sell = ?", price).Find(&b).Error
	if err != nil {
		return nil
	}
	err = db.Table(fishTable).Where("sell = ?", price).Find(&f).Error
	if err != nil {
		return nil
	}
	return &newhorizons.Combined{
		Bugs:   b,
		Fishes: f,
	}
}

// findByName retrieves a bug database entry by a given item name
//
// If none is found, it will return nil
func findBugByName(item, tablename string, db *gorm.DB) *newhorizons.BugEntry {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var entry newhorizons.BugEntry
	err := db.Table(tablename).Where("name = ?", item).First(&entry).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return &entry
}

// findByName retrieves a fish database entry by a given item name
//
// If none is found, it will return nil
func findFishByName(item, tablename string, db *gorm.DB) *newhorizons.FishEntry {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var entry newhorizons.FishEntry
	err := db.Table(tablename).Where("name = ?", item).First(&entry).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return &entry
}

// findAllByHemisphere returns an object which contains all bug and fish
// entries depending on hemisphere
func findAllByHemisphere(bugTable, fishTable string, db *gorm.DB) *newhorizons.Combined {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var b []newhorizons.BugEntry
	var f []newhorizons.FishEntry
	err := db.Raw("SELECT * FROM " + bugTable).Find(&b).Error
	if err != nil {
		return nil
	}
	err = db.Raw("SELECT * FROM " + fishTable).Find(&f).Error
	if err != nil {
		return nil
	}
	return &newhorizons.Combined{
		Bugs:   b,
		Fishes: f,
	}
}
