package testrail

import (
	"strconv"
)

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

// PaginatedSuites represents a Test Suites response
// from the TestRail API (suites w/ pagination)
type PaginatedSuites struct {
	Offset int             `json:"offset"`
	Limit  int             `json:"limit"`
	Size   int             `json:"size"`
	Suites []Suite         `json:"suites"`
	Links  PaginationLinks `json:"_links"`
}

type PaginationLinks struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

// GetSuite returns the suite suiteID
func (c *Client) GetSuite(suiteID int) (Suite, error) {
	returnSuite := Suite{}
	err := c.sendRequest("GET", "get_suite/"+strconv.Itoa(suiteID), nil, &returnSuite)
	return returnSuite, err
}

// GetSuites returns the list of suites on project projectID.
// Pagination is handled internally (calling may result in 
// multiple API requests).
func (c *Client) GetSuites(projectID int) ([]Suite, error) {
	returnSuites := []Suite{}
	url := "get_suites/" + strconv.Itoa(projectID)
	for {
		paginated := PaginatedSuites{}
		err := c.sendRequest("GET", url, nil, &paginated)
		if err != nil {
			return nil, err
		}
		returnSuites = append(returnSuites, paginated.Suites...)
		if paginated.Links.Next == "" {
			return returnSuites, nil
		}
		url = paginated.Links.Next
	}
}

// AddSuite creates a new suite on projectID and returns it
func (c *Client) AddSuite(projectID int, newSuite SendableSuite) (Suite, error) {
	createdSuite := Suite{}
	err := c.sendRequest("POST", "add_suite/"+strconv.Itoa(projectID), newSuite, &createdSuite)
	return createdSuite, err
}

// UpdateSuite updates the suite suiteID and returns it
func (c *Client) UpdateSuite(suiteID int, update SendableSuite) (Suite, error) {
	updatedSuite := Suite{}
	err := c.sendRequest("POST", "update_suite/"+strconv.Itoa(suiteID), update, &updatedSuite)
	return updatedSuite, err
}

// DeleteSuite delete the suite suiteID
func (c *Client) DeleteSuite(suiteID int) error {
	return c.sendRequest("POST", "delete_suite/"+strconv.Itoa(suiteID), nil, nil)
}
