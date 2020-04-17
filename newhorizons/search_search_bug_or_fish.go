package newhorizons

// SearchBugOrFishResult represents all possible returned fields
// of a fish or bug entry in the database
type SearchBugOrFishResult struct {
	Num          int    `json:"num,omitempty"`
	Name         string `json:"name,omitempty"`
	Image        string `json:"image,omitempty"`
	House        string `json:"house,omitempty"`
	Sell         int    `json:"sell,omitempty"`
	Where        string `json:"where,omitempty"`
	Weather      string `json:"weather,omitempty"`
	Rarity       string `json:"rarity,omitempty"`
	StartTime    string `json:"start_time,omitempty"`
	EndTime      string `json:"end_time,omitempty"`
	Jan          string `json:"jan,omitempty"`
	Feb          string `json:"feb,omitempty"`
	Mar          string `json:"mar,omitempty"`
	Apr          string `json:"apr,omitempty"`
	May          string `json:"may,omitempty"`
	Jun          string `json:"jun,omitempty"`
	Jul          string `json:"jul,omitempty"`
	Aug          string `json:"aug,omitempty"`
	Sep          string `json:"sep,omitempty"`
	Oct          string `json:"oct,omitempty"`
	Nov          string `json:"nov,omitempty"`
	Dec          string `json:"dec,omitempty"`
	Color1       string `json:"color_1,omitempty"`
	Color2       string `json:"color_2,omitempty"`
	ItemFilename string `json:"item_filename,omitempty"`
	InternalID   int    `json:"internal_id,omitempty"`
}
