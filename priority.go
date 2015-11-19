package testrail

type Priorities struct {
	ID        int    `json:"id"`
	IsDefault bool   `json:"is_default"`
	Name      string `json:"name"`
	Priority  int    `json:"priority"`
	ShortName string `json:"short_name"`
}

// Returns a list of available priorities.
func (c *Client) GetPriorities() ([]Priorities, error) {
	prios := []Priorities{}
	err := c.sendRequest("GET", "get_priorities/", nil, &prios)
	return prios, err
}
