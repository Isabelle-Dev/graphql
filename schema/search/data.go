package search

import "github.com/jinzhu/gorm"

// Entry represents a database entry of either an insect or
// fish in the postgres database
type Entry struct {
	Name        string `gorm:"not null"`
	SellPrice   int    `gorm:"column:sell_price"`
	NorthSt     string `gorm:"type:varchar(255);column:north_start"`
	NorthEnd    string `gorm:"type:varchar(255);column:north_end"`
	NorthMonths string `gorm:"type:varchar(255);column:north_hemi_months"`
	SouthSt     string `gorm:"type:varchar(255);column:south_start"`
	SouthEnd    string `gorm:"type:varchar(255);column:south_end"`
	SouthMonths string `gorm:"type:varchar(255);column:south_hemi_months"`
	Time        string `gorm:"type:varchar(255);column:time_of_day"`
	Location    string `gorm:"type:varchar(255);column:location"`
	Image       string `gorm:"type:varchar(255);column:image"`
	Type        string `gorm:"type:varchar(255);column:type"`
}

// findByName retrieves a database entry by a given item name
//
// If none is found, it will return an error
func findByName(item, tablename string, db *gorm.DB) *Entry {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var entry Entry
	err := db.Table(tablename).Where("name = ?", item).First(&entry).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return &entry
}
