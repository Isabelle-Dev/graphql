package newhorizons

// FenceEntry represents a fence object from the database
type FenceEntry struct {
	Name        string
	Image       string
	DIY         string
	Buy         int
	Sell        int
	Source      string
	SourceNotes string
}
