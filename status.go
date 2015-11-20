package testrail

type Statuses struct {
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

// Return the list of all possible statuses
func (c *Client) GetStatuses() ([]Statuses, error) {
	returnStatuses := []Statuses{}
	err := c.sendRequest("GET", "get_statuses/", nil, &returnStatuses)
	return returnStatuses, err
}
