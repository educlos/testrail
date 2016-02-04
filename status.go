package testrail

// Status represents a Status
type Status struct {
	ColorBright int    `json:"color_bright"`
	ColorDark   int    `json:"color_dark"`
	ColorMedium int    `json:"color_medium"`
	ID          int    `json:"id"`
	IsFinal     bool   `json:"is_final"`
	IsSystem    bool   `json:"is_system"`
	IsUntested  bool   `json:"is_untested"`
	Label       string `json:"label"`
	Name        string `json:"name"`
}

// GetStatuses return the list of all possible statuses
func (c *Client) GetStatuses() (statuses []Status, err error) {
	err = c.sendRequest("GET", "get_statuses", nil, &statuses)
	return
}
