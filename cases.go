package testrail

import "strconv"

type Case struct {
	CreatedBy            int          `json:"created_by"`
	CreatedOn            int          `json:"created_on"`
	CustomExpected       string       `json:"custom_expected"`
	CustomPreconds       string       `json:"custom_preconds"`
	CustomSteps          string       `json:"custom_steps"`
	CustomStepsSeparated []CustomStep `json:"custom_steps_separated"`
	Estimate             string       `json:"estimate"`
	EstimateForecast     string       `json:"estimate_forecast"`
	ID                   int          `json"id"`
	MilestoneId          int          `json:"milestone_id"`
	PriorityId           int          `json:"priority_id"`
	Refs                 string       `json:"refs"`
	SectionId            int          `json:"section_id"`
	SuiteId              int          `json:"suite_id"`
	Title                string       `json:"title"`
	TypeId               int          `json:"type_id"`
	UpdatedBy            int          `json:"updated_by"`
	UdpatedOn            int          `json:"updated_on"`
}

type CustomStep struct {
	Content  string `json:"content"`
	Expected string `json:"expected"`
}

type RequestFilter struct {
	CreatedAfter  string `json:"created_after"`
	CreatedBefore string `json:"created_before"`
	CreatedBy     []int  `json:"created_by"`
	MilestoneId   []int  `json:"milestone_id"`
	PriorityId    []int  `json:"priority_id"`
	TypeId        []int  `json:"type_id"`
	UpdatedAfter  string `json:"updated_after"`
	UpdatedBefore string `json:"updated_before"`
	UpdatedBy     []int  `json:"updated_by"`
}

type SendableCase struct {
	Title       string       `json:"title"`
	TypeId      int          `json:"type_id"`
	PriorityId  int          `json:"priority_id"`
	Estimate    string       `json:"estimate"`
	MilestoneId int          `json:"milestone_id"`
	Refs        string       `json:"refs"`
	Checkbox    bool         `json:"custom_checkbox"`
	Date        string       `json:"custom_date"`
	Dropdown    int          `json:"custom_dropdown"`
	Integer     int          `json:"custom_integer"`
	Milestone   int          `json:"custom_milestone"`
	MultiSelect []int        `json:"custom_multi-select"`
	Steps       []CustomStep `json:"custom_steps"`
	String      string       `json:"custom_string"`
	Text        string       `json:"custom_text`
	URL         string       `json:"custom_url`
	User        int          `json:"custom_user"`
}

// Returns the existing test case caseID
func (c *Client) GetCase(caseID int) (Case, error) {
	returnCase := Case{}
	err := c.sendRequest("GET", "get_case/"+strconv.Itoa(caseID), nil, &returnCase)
	return returnCase, err
}

// Returns a list of test cases for a test suite or specific section in a test suite.
func (c *Client) GetCases(projectID, suiteID int, sectionID ...int) ([]Case, error) {
	uri := "get_cases/" + strconv.Itoa(projectID) + "&suite_id=" + strconv.Itoa(suiteID)
	if len(sectionID) > 0 {
		uri = uri + "&section_id=" + strconv.Itoa(sectionID[0])
	}

	returnCases := []Case{}
	err := c.sendRequest("GET", uri, nil, &returnCases)
	return returnCases, err
}

// Creates a new test case and return the created test case
func (c *Client) AddCase(sectionID int, newCase SendableCase) (Case, error) {
	createdCase := Case{}
	err := c.sendRequest("POST", "add_case/"+strconv.Itoa(sectionID), newCase, &createdCase)
	return createdCase, err
}

// Updates the existing test case caseID
func (c *Client) UpdateCase(caseID int, updates SendableCase) (Case, error) {
	updatedCase := Case{}
	err := c.sendRequest("POST", "update_case/"+strconv.Itoa(caseID), updates, &updatedCase)
	return updatedCase, err
}

// Deletes the existing test case caseID
func (c *Client) DeleteCase(caseID int) error {
	return c.sendRequest("POST", "delete_case/"+strconv.Itoa(caseID), nil, nil)
}
