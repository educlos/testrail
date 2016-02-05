package testrail

import "fmt"

// Suite represenst a Test Suite
type Suite struct {
	CompletedOn int    `json:"completed_on"`
	Description string `json:"description"`
	ID          int    `json:"id"`
	IsBaseline  bool   `json:"is_baseline"`
	IsCompleted bool   `json:"is_completed"`
	IsMaster    bool   `json:"is_master"`
	Name        string `json:"name"`
	ProjectID   int    `json:"project_id"`
	URL         string `json:"url"`
}

// SendableSuite represents a Test Suite
// that can be created or updated via the api
type SendableSuite struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// GetSuite returns the suite suiteID
func (c *Client) GetSuite(suiteID int) (suite Suite, err error) {
	err = c.sendRequest("GET", fmt.Sprintf("get_suite/%d", suiteID), nil, &suite)
	return
}

// GetSuites returns the list of suites on project projectID
func (c *Client) GetSuites(projectID int) (suites []Suite, err error) {
	err = c.sendRequest("GET", fmt.Sprintf("get_suites/%d", projectID), nil, &suites)
	return
}

// AddSuite creates a new suite on projectID and returns it
func (c *Client) AddSuite(projectID int, newSuite SendableSuite) (suite Suite, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("add_suite/%d", projectID), newSuite, &suite)
	return
}

// UpdateSuite updates the suite suiteID and returns it
func (c *Client) UpdateSuite(suiteID int, update SendableSuite) (suite Suite, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("update_suite/%d", suiteID), update, &suite)
	return
}

// DeleteSuite delete the suite suiteID
func (c *Client) DeleteSuite(suiteID int) error {
	return c.sendRequest("POST", fmt.Sprintf("delete_suite/%d", suiteID), nil, nil)
}
