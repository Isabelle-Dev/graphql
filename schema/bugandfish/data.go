package bugandfish

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func findByShadow(shadow, fishTable string, db *gorm.DB) []interface{} {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var f []newhorizons.FishEntry
	err := db.Table(fishTable).Where("shadow = ?", shadow).Find(&f).Error
	if err != nil {
		return nil
	}
	var ret []interface{}
	for _, e := range f {
		entry := findFishByName(e.Name, "north_fish", "south_fish", db)
		if et, ok := entry.(*newhorizons.Fish); ok {
			ret = append(ret, e.ToGraphQL(et.NorthernHemi.Months, et.SouthernHemi.Months))
		}
	}
	return ret
}

// findByRarity returns all bugs and fishes for a given rarity
func findByRarity(rarity, bugTable, fishTable string, db *gorm.DB) *newhorizons.Combined {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var b []newhorizons.BugEntry
	var f []newhorizons.FishEntry
	var bug []interface{}
	var fish []interface{}
	err := db.Table(bugTable).Where("rarity = ?", rarity).Find(&b).Error
	if err != nil {
		return nil
	}
	err = db.Table(fishTable).Where("rarity = ?", rarity).Find(&f).Error
	if err != nil {
		return nil
	}
	for i := 0; i < len(b); i++ {
		s := findBugByName(b[i].Name, "north_bug", "south_bug", db)
		if e, ok := s.(*newhorizons.Bug); ok {
			bug = append(bug, b[i].ToGraphQL(e.NorthernHemi.Months, e.SouthernHemi.Months))
		}
	}
	for i := 0; i < len(f); i++ {
		s := findFishByName(f[i].Name, "north_fish", "south_fish", db)
		if e, ok := s.(*newhorizons.Fish); ok {
			fish = append(fish, f[i].ToGraphQL(e.NorthernHemi.Months, e.SouthernHemi.Months))
		}
	}
	return &newhorizons.Combined{
		Bugs:   bug,
		Fishes: fish,
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
	var bug []interface{}
	var fish []interface{}
	for i := 0; i < len(b); i++ {
		s := findBugByName(b[i].Name, "north_bug", "south_bug", db)
		if e, ok := s.(*newhorizons.Bug); ok {
			bug = append(bug, b[i].ToGraphQL(e.NorthernHemi.Months, e.SouthernHemi.Months))
		}
	}
	for i := 0; i < len(f); i++ {
		s := findFishByName(f[i].Name, "north_fish", "south_fish", db)
		if e, ok := s.(*newhorizons.Fish); ok {
			fish = append(fish, f[i].ToGraphQL(e.NorthernHemi.Months, e.SouthernHemi.Months))
		}
	}
	return &newhorizons.Combined{
		Bugs:   bug,
		Fishes: fish,
	}
}

// findByPrice returns the object of all bug and fish listings that match a given price
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
	var bug []interface{}
	var fish []interface{}
	for i := 0; i < len(b); i++ {
		s := findBugByName(b[i].Name, "north_bug", "south_bug", db)
		if e, ok := s.(*newhorizons.Bug); ok {
			bug = append(bug, b[i].ToGraphQL(e.NorthernHemi.Months, e.SouthernHemi.Months))
		}
	}
	for i := 0; i < len(f); i++ {
		s := findFishByName(f[i].Name, "north_fish", "south_fish", db)
		if e, ok := s.(*newhorizons.Fish); ok {
			fish = append(fish, f[i].ToGraphQL(e.NorthernHemi.Months, e.SouthernHemi.Months))
		}
	}
	return &newhorizons.Combined{
		Bugs:   bug,
		Fishes: fish,
	}
}

// findByName retrieves a bug database entry by a given item name
//
// If none is found, it will return nil
func findBugByName(item, northTable, southTable string, db *gorm.DB) interface{} {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var entry newhorizons.BugEntry
	err := db.Table(northTable).Where("name = ?", item).First(&entry).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	northMonths := extractMonths(entry)
	err = db.Table(southTable).Where("name = ?", item).First(&entry).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	southMonths := extractMonths(entry)
	return entry.ToGraphQL(northMonths, southMonths)
}

// findByName retrieves a fish database entry by a given item name
//
// If none is found, it will return nil
func findFishByName(item, northTable, southTable string, db *gorm.DB) interface{} {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var entry newhorizons.FishEntry
	err := db.Table(northTable).Where("name = ?", item).First(&entry).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	northMonths := extractMonths(entry)
	err = db.Table(southTable).Where("name = ?", item).First(&entry).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	southMonths := extractMonths(entry)
	return entry.ToGraphQL(northMonths, southMonths)
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
	var bug []interface{}
	var fish []interface{}
	for i := 0; i < len(b); i++ {
		s := findBugByName(b[i].Name, "north_bug", "south_bug", db)
		if e, ok := s.(*newhorizons.Bug); ok {
			bug = append(bug, b[i].ToGraphQL(e.NorthernHemi.Months, e.SouthernHemi.Months))
		}
	}
	for i := 0; i < len(f); i++ {
		s := findFishByName(f[i].Name, "north_fish", "south_fish", db)
		if e, ok := s.(*newhorizons.Fish); ok {
			fish = append(fish, f[i].ToGraphQL(e.NorthernHemi.Months, e.SouthernHemi.Months))
		}
	}
	return &newhorizons.Combined{
		Bugs:   bug,
		Fishes: fish,
	}
}
