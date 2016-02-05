package testrail

type CaseType struct {
	ID        int    `json:"id"`
	IsDefault bool   `json:"is_default"`
	Name      string `json:"name"`
}

// Returns a list of available test case types
func (c *Client) GetCaseTypes() (types []CaseType, err error) {
	err = c.sendRequest("GET", "get_case_types", nil, &types)
	return
}
