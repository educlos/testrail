package testrail

import (
	"fmt"
	"net/url"
)

// Case represents a Test Case
type Case struct {
	CreatedBy            int          `json:"created_by"`
	CreatedOn            int          `json:"created_on"`
	CustomExpected       string       `json:"custom_expected"`
	CustomPreconds       string       `json:"custom_preconds"`
	CustomSteps          string       `json:"custom_steps"`
	CustomStepsSeparated []CustomStep `json:"custom_steps_separated"`
	Estimate             string       `json:"estimate"`
	EstimateForecast     string       `json:"estimate_forecast"`
	ID                   int          `json:"id"`
	MilestoneID          int          `json:"milestone_id"`
	PriorityID           int          `json:"priority_id"`
	Refs                 string       `json:"refs"`
	SectionID            int          `json:"section_id"`
	SuiteID              int          `json:"suite_id"`
	Title                string       `json:"title"`
	TypeID               int          `json:"type_id"`
	UpdatedBy            int          `json:"updated_by"`
	UdpatedOn            int          `json:"updated_on"`
}

// CustomStep represents the custom steps
// a Test Case can have
type CustomStep struct {
	Content  string `json:"content"`
	Expected string `json:"expected"`
}

// RequestFilterForCases represents the filters
// usable to get the test cases
type RequestFilterForCases struct {
	CreatedAfter  string  `json:"created_after,omitempty"`
	CreatedBefore string  `json:"created_before,omitempty"`
	CreatedBy     IntList `json:"created_by,omitempty"`
	MilestoneID   IntList `json:"milestone_id,omitempty"`
	PriorityID    IntList `json:"priority_id,omitempty"`
	TypeID        IntList `json:"type_id,omitempty"`
	UpdatedAfter  string  `json:"updated_after,omitempty"`
	UpdatedBefore string  `json:"updated_before,omitempty"`
	UpdatedBy     IntList `json:"updated_by,omitempty"`
}

// SendableCase represents a Test Case
// that can be created or updated via the api
type SendableCase struct {
	Title       string       `json:"title"`
	TypeID      int          `json:"type_id,omitempty"`
	PriorityID  int          `json:"priority_id,omitempty"`
	Estimate    string       `json:"estimate,omitempty"`
	MilestoneID int          `json:"milestone_id,omitempty"`
	Refs        string       `json:"refs,omitempty"`
	Checkbox    bool         `json:"custom_checkbox,omitempty"`
	Date        string       `json:"custom_date,omitempty"`
	Dropdown    int          `json:"custom_dropdown,omitempty"`
	Integer     int          `json:"custom_integer,omitempty"`
	Milestone   int          `json:"custom_milestone,omitempty"`
	MultiSelect []int        `json:"custom_multi_select,omitempty"`
	Steps       []CustomStep `json:"custom_steps,omitempty"`
	String      string       `json:"custom_string,omitempty"`
	Text        string       `json:"custom_text,omitempty"`
	URL         string       `json:"custom_url,omitempty"`
	User        int          `json:"custom_user,omitempty"`
}

// GetCase returns the existing Test Case caseID
func (c *Client) GetCase(caseID int) (case_ Case, err error) {
	err = c.sendRequest("GET", fmt.Sprintf("get_case/%d", caseID), nil, &case_)
	return
}

// GetCases returns a list of Test Cases on project projectID
// for a Test Suite suiteID
// or for specific section sectionID in a Test Suite
func (c *Client) GetCases(projectID, suiteID int, sectionID ...int) (cases []Case, err error) {
	vals := make(url.Values)
	vals.Set("suite_id", fmt.Sprintf("%d", suiteID))

	if len(sectionID) > 0 {
		vals.Set("section_id", fmt.Sprintf("%d", sectionID[0]))
	}

	err = c.sendRequest("GET", fmt.Sprintf("get_cases/%d?%s", projectID, vals.Encode()), nil, &cases)
	return
}

// GetCasesWithFilters returns a list of Test Cases on project projectID
// for a Test Suite suiteID validating the filters
func (c *Client) GetCasesWithFilters(projectID, suiteID int, filters ...RequestFilterForCases) (cases []Case, err error) {
	vals := make(url.Values)
	vals.Set("suite_id", fmt.Sprintf("%d", suiteID))

	loadOptionalFilters(vals, filters)

	err = c.sendRequest("GET", fmt.Sprintf("get_cases/%d?%s", projectID, vals.Encode()), nil, &cases)
	return
}

// AddCase creates a new Test Case newCase and returns it
func (c *Client) AddCase(sectionID int, newCase SendableCase) (case_ Case, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("add_case/%d", sectionID), newCase, &case_)
	return
}

// UpdateCase updates an existing Test Case caseID and returns it
func (c *Client) UpdateCase(caseID int, updates SendableCase) (case_ Case, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("update_case/%d", caseID), updates, &case_)
	return
}

// DeleteCase deletes the existing Test Case caseID
func (c *Client) DeleteCase(caseID int) error {
	return c.sendRequest("POST", fmt.Sprintf("delete_case/%d", caseID), nil, nil)
}
