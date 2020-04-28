package tool

import (
	"github.com/Isabelle-Dev/isabelle-graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

func execute(dbStr string, db *gorm.DB) []*newhorizons.Tool {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var tools []newhorizons.ToolEntry
	db.Raw(dbStr).Find(&tools)
	if len(tools) == 0 {
		return nil
	}
	return toToolSlice(tools, db)
}

func findByName(name, tablename string, db *gorm.DB) *newhorizons.Tool {
	db.RWMutex.RLock()
	defer db.RWMutex.RUnlock()
	var tool []newhorizons.ToolEntry
	db.Table(tablename).Where("name = ?", name).Find(&tool)
	if len(tool) == 0 {
		return nil
	}
	v, s := extractVS(tool)
	return tool[0].ToGraphQL(v, s)
}
