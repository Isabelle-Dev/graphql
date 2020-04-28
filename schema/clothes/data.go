package clothes

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
func execute(dbStr string, db *gorm.DB) []*newhorizons.Clothes {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var clothes []newhorizons.ClothesEntry
	db.Raw(dbStr).Find(&clothes)
	if len(clothes) == 0 {
		return nil
	}
	return toClothesSlice(clothes, db)
}

func findByName(name, tablename string, db *gorm.DB) *newhorizons.Clothes {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var clothes []newhorizons.ClothesEntry
	db.Table(tablename).Where("name = ?", name).Find(&clothes)
	if len(clothes) == 0 {
		return nil
	}
	v, t := extractVT(clothes)
	return clothes[0].ToGraphQL(v, t)
}
