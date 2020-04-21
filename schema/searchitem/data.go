package searchitem

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// findByName retrieves a list of entries from the item table by
// a given name
func findByName(name, tablename string, db *gorm.DB) *newhorizons.ModifiedItem {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var items []newhorizons.Item
	err := db.Table(tablename).Where("name = ?", name).Find(&items).Error
	if err != nil {
		return nil
	}
	var variants []string
	var patterns []string
	var images []string
	for _, entry := range items {
		if entry.PatternTitle != "NA" {
			patterns = append(patterns, entry.Pattern)
		}
		if entry.Variation != "NA" {
			variants = append(variants, entry.Variation)
		}
		images = append(images, entry.Image)
	}
	return toModifiedItem(items[0], variants, patterns, images)
}

// helper func which turns an item into a modified item
func toModifiedItem(item newhorizons.Item, v, p, i []string) *newhorizons.ModifiedItem {
	return &newhorizons.ModifiedItem{
		Name:             item.Name,
		Image:            i,
		Variation:        v,
		BodyTitle:        item.BodyTitle,
		Pattern:          p,
		PatternTitle:     item.PatternTitle,
		DIY:              item.DIY,
		BodyCustomize:    item.BodyCustomize,
		PatternCustomize: item.PatternCustomize,
		KitCost:          item.KitCost,
		Buy:              item.Buy,
		Sell:             item.Sell,
		Color1:           item.Color1,
		Color2:           item.Color2,
		Size:             item.Size,
		Source:           item.Source,
		SourceNotes:      item.SourceNotes,
		Version:          item.Version,
		HHAConcept1:      item.HHAConcept1,
		HHAConcept2:      item.HHAConcept2,
		HHASeries:        item.HHASeries,
		HHASet:           item.HHASet,
		Interact:         item.Interact,
		Tag:              item.Tag,
		SpeakerType:      item.SpeakerType,
		LightingType:     item.LightingType,
		Catalog:          item.Catalog,
		InternalID:       item.InternalID,
	}
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
