package tool

import (
	"github.com/Isabelle-Dev/graphql/newhorizons"
	"github.com/jinzhu/gorm"
)

// execute runs a raw SQL query string and returns a slice of
// newhorizons type which match the query criteria
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

// findByName retrieves an entry from the tool table by
// a given name
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
