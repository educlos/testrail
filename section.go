package testrail

import (
	"fmt"
	"net/url"
)

// Section represents a Test Suite Section
type Section struct {
	Depth        int    `json:"depth"`
	Description  string `json:"description"`
	DisplayOrder int    `json:"display_order"`
	ID           int    `json:"id"`
	ParentID     int    `json:"parent_id"`
	Name         string `json:"name"`
	SuiteID      int    `json:"suite_id"`
}

// SendableSection represents a Test Suite Section
// that can be created via the api
type SendableSection struct {
	Description string `json:"description,omitempty"`
	SuiteID     int    `json:"suite_id,omitempty"`
	ParentID    int    `json:"parent_id,omitempty"`
	Name        string `json:"name"`
}

// UpdatableSection represents a Test Suite Section
// that can be updated via the api
type UpdatableSection struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

// GetSection returns the section sectionID
func (c *Client) GetSection(sectionID int) (section Section, err error) {
	err = c.sendRequest("GET", fmt.Sprintf("get_section/%d", sectionID), nil, &section)
	return
}

// GetSections returns the list of sections of projectID
// present in suite suiteID, if specified
func (c *Client) GetSections(projectID int, suiteID ...int) (sections []Section, err error) {
	vals := make(url.Values)

	if len(suiteID) > 0 {
		vals.Set("suite_id", fmt.Sprintf("%d", suiteID[0]))
	}

	err = c.sendRequest("GET", fmt.Sprintf("get_sections/%d?%s", projectID, vals.Encode()), nil, &sections)
	return
}

// AddSection creates a new section on projectID and returns it
func (c *Client) AddSection(projectID int, newSection SendableSection) (section Section, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("add_section/%d", projectID), newSection, &section)
	return
}

// UpdateSection updates the section sectionID and returns it
func (c *Client) UpdateSection(sectionID int, update UpdatableSection) (section Section, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("update_section/%d", sectionID), update, &section)
	return
}

// DeleteSection deletes the section sectionID
func (c *Client) DeleteSection(sectionID int) error {
	return c.sendRequest("POST", fmt.Sprintf("delete_section/%d", sectionID), nil, nil)
}
