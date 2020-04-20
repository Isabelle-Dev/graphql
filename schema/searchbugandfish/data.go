package searchbugandfish

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func findByShadow(shadow, fishTable string, db *gorm.DB) []newhorizons.FishTimeAgnostic {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var f []newhorizons.FishTimeAgnostic
	err := db.Table(fishTable).Where("shadow = ?", shadow).Find(&f).Error
	if err != nil {
		return nil
	}
	return f
}

// findByRarity returns all bugs and fishes for a given rarity
func findByRarity(rarity, bugTable, fishTable string, db *gorm.DB) *newhorizons.CombinedAgnostic {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var b []newhorizons.BugTimeAgnostic
	var f []newhorizons.FishTimeAgnostic
	err := db.Table(bugTable).Where("rarity = ?", rarity).Find(&b).Error
	if err != nil {
		return nil
	}
	err = db.Table(fishTable).Where("rarity = ?", rarity).Find(&f).Error
	if err != nil {
		return nil
	}
	return &newhorizons.CombinedAgnostic{
		Bugs:   b,
		Fishes: f,
	}
}

// findByMonth will return an object with all bugs and fishes available in a given month
func findByMonth(month, bugTable, fishTable string, db *gorm.DB) *newhorizons.Combined {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var b []newhorizons.BugEntry
	var f []newhorizons.FishEntry
	err := db.Table(bugTable).Where(month+" = ?", "Yes").Find(&b).Error
	if err != nil {
		return nil

	}
	err = db.Table(fishTable).Where(month+" = ?", "Yes").Find(&f).Error
	if err != nil {
		return nil
	}
	return &newhorizons.Combined{
		Bugs:   b,
		Fishes: f,
	}
}

// findByPrice returns the object of all bug and fish listings that match a given price
func findByPrice(price int, bugTable, fishTable string, db *gorm.DB) *newhorizons.CombinedAgnostic {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var b []newhorizons.BugTimeAgnostic
	var f []newhorizons.FishTimeAgnostic
	err := db.Table(bugTable).Where("sell = ?", price).Find(&b).Error
	if err != nil {
		return nil
	}
	err = db.Table(fishTable).Where("sell = ?", price).Find(&f).Error
	if err != nil {
		return nil
	}
	return &newhorizons.CombinedAgnostic{
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
