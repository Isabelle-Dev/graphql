package newhorizons

type SearchBugOrFishResult struct {
	Name      string `json:"name,omitempty"`
	SellPrice int    `json:"sell_price,omitempty"`
	Location  string `json:"location,omitempty"`
	Time      string `json:"time,omitempty"`
	Type      string `json:"type,omitempty"`
}
