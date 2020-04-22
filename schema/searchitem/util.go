package searchitem

import "github.com/Isabelle-Dev/isabelle-graphql/newhorizons"

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

func extractVPI(item []newhorizons.Item) ([]string, []string, []string) {
	var variants []string
	var patterns []string
	var images []string
	for _, entry := range item {
		if entry.PatternTitle != "NA" {
			patterns = append(patterns, entry.Pattern)
		}
		if entry.Variation != "NA" {
			variants = append(variants, entry.Variation)
		}
		images = append(images, entry.Image)
	}
	return variants, patterns, images
}
