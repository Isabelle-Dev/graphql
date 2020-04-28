package newhorizons

// Tool represents a tool entry graphql is able to parse and display
type Tool struct {
	Name        string
	Variant     []ToolVariant
	BodyTitle   string
	DIY         string
	Customize   string
	KitCost     int
	Uses        int
	Buy         int
	Sell        int
	Set         string
	Source      string
	SourceNotes string
}

// ToolVariant represents a tool's variation parameters
type ToolVariant struct {
	Image     string
	Variation string
}

// ToolEntry represents a tool in the database
type ToolEntry struct {
	Name        string
	Image       string
	Variation   string
	BodyTitle   string `gorm:"column:bodytitle"`
	DIY         string
	Customize   string
	KitCost     int `gorm:"column:kitcost"`
	Uses        int
	Buy         int
	Sell        int
	Set         string
	Source      string
	SourceNotes string `gorm:"column:sourcenotes"`
}

// ToGraphQL (ToolEntry) converts a tool entry into a tool object
func (te ToolEntry) ToGraphQL(t []ToolVariant) *Tool {
	return &Tool{
		Name:        te.Name,
		Variant:     t,
		BodyTitle:   te.BodyTitle,
		DIY:         te.DIY,
		Customize:   te.Customize,
		KitCost:     te.KitCost,
		Uses:        te.Uses,
		Buy:         te.Buy,
		Sell:        te.Sell,
		Set:         te.Set,
		Source:      te.Source,
		SourceNotes: te.SourceNotes,
	}
}
