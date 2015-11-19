package testrail

import "strconv"

type Result struct {
	AssignedtoID      int                `json:"assignedto_id"`
	Comment           string             `json:"comment"`
	CreatedBy         int                `json:"created_by"`
	CreatedOn         int                `json:"created_on"`
	CustomStepResults []CustomStepResult `json:"custom_step_results"`
	Defects           string             `json:"defects"`
	Elapsed           int                `json:"elapsed"`
	ID                int                `json:"id"`
	StatusID          int                `json:"status_id"`
	TestID            int                `json:"test_id"`
	Version           string             `json:"version"`
}

type CustomComments struct {
	Checkbox    bool               `json:"checkbox"`
	Date        string             `json:"date"`
	Dropdown    int                `json:"dropdown"`
	Integer     int                `json:"integer"`
	Milestone   int                `json:"milestone"`
	MultiSelect []int              `json:"multi-select"`
	StepResults []CustomStepResult `json:"step_results"`
	String      string             `json:"string"`
	Text        string             `json:"text"`
	URL         string             `json:"url"`
	User        int                `json:"user"`
}

type CustomStepResult struct {
	Content  string `json:"content"`
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
	StatusID int    `json:"status_id"`
}

type RequestFilterForCaseResults struct {
	Limit    *int  `json:"limit,omitempty"`
	Offest   *int  `json:"offset, omitempty"`
	StatusID []int `json:"status_id,omitempty"`
}

type RequestFilterForRunResults struct {
	CreatedAfter  string `json:"created_after,omitempty"`
	CreatedBefore string `json:"created_before,omitempty"`
	CreatedBy     []int  `json:"created_by,omitempty"`
	Limit         *int   `json:"limit,omitempty"`
	Offest        *int   `json:"offset, omitempty"`
	StatusID      []int  `json:"status_id,omitempty"`
}

type SendableResult struct {
	StatusID     int                `json:"status_id,omitempty"`
	Comment      string             `json:"comment,omitempty"`
	Version      string             `json:"version,omitempty"`
	Elapsed      string             `json:"elapsed,omitempty"`
	Defects      string             `json:"defects,omitempty"`
	AssignedToID int                `json:"assignedto_id,omitempty"`
	Checkbox     bool               `json:"custom_checkbox,omitempty"`
	Date         string             `json:"custom_date,omitempty"`
	Dropdown     int                `json:"custom_dropdown,omitempty"`
	Integer      int                `json:"custom_integer,omitempty"`
	Milestone    int                `json:"custom_milestone,omitempty"`
	MultiSelect  []int              `json:"custom_multi-select,omitempty"`
	StepsResults []CustomStepResult `json:"custom_step_results,omitempty"`
	String       string             `json:"custom_string,omitempty"`
	Text         string             `json:"custom_text,omitempty"`
	URL          string             `json:"custom_url,omitempty"`
	User         int                `json:"custom_user,omitempty"`
}

type SendableResults struct {
	Results []Results `json:"results"`
}

type Results struct {
	TestID int `json:"test_id"`
	SendableResult
}

type SendableResultsForCase struct {
	Results []Results `json:"results"`
}

type ResultsForCase struct {
	CaseID int `json:"case_id"`
	SendableResult
}

// Returns a list of test results for the test testID
func (c *Client) GetResults(testID int, filters ...RequestFilterForCaseResults) ([]Result, error) {
	returnResults := []Result{}
	uri := "get_results/" + strconv.Itoa(testID)

	if len(filters) > 0 {
		uri = applyFiltersForCaseResults(uri, filters[0])
	}
	err := c.sendRequest("GET", uri, nil, &returnResults)
	return returnResults, err
}

// Returns a list of test results for the case caseID on run runID
func (c *Client) GetResultsForCase(runID, caseID int, filters ...RequestFilterForCaseResults) ([]Result, error) {
	returnResults := []Result{}
	uri := "get_results_for_case/" + strconv.Itoa(runID) + "/" + strconv.Itoa(caseID)

	if len(filters) > 0 {
		uri = applyFiltersForCaseResults(uri, filters[0])
	}
	err := c.sendRequest("GET", uri, nil, &returnResults)
	return returnResults, err
}

// Returns a list of test results for the run runID
func (c *Client) GetResultsForRun(runID int, filters ...RequestFilterForRunResults) ([]Result, error) {
	returnResults := []Result{}
	uri := "get_results_for_run/" + strconv.Itoa(runID)

	if len(filters) > 0 {
		uri = applyFiltersForRunResults(uri, filters[0])
	}
	err := c.sendRequest("GET", uri, nil, &returnResults)
	return returnResults, err
}

// Adds a new test result, comment or assigns a test to testID
func (c *Client) AddResult(testID int, newResult SendableResult) (Result, error) {
	createdResult := Result{}
	err := c.sendRequest("POST", "add_result/"+strconv.Itoa(testID), newResult, &createdResult)
	return createdResult, err
}

// Adds a new test result, comment or assigns a test to the case caseID on run runID
func (c *Client) AddResultForCase(runID, caseID int, newResult SendableResult) (Result, error) {
	createdResult := Result{}
	uri := "add_result_for_case/" + strconv.Itoa(runID) + "/" + strconv.Itoa(caseID)
	err := c.sendRequest("POST", uri, newResult, &createdResult)
	return createdResult, err
}

// Adds a new test result, comment or assigns a test to runID
func (c *Client) AddResults(runID int, newResult SendableResults) ([]Result, error) {
	createdResult := []Result{}
	err := c.sendRequest("POST", "add_results/"+strconv.Itoa(runID), newResult, &createdResult)
	return createdResult, err
}

// Adds one or more new test results, comments or assigns one or more tests to run runID
func (c *Client) AddResultsForCase(runID int, newResult SendableResultsForCase) (Result, error) {
	createdResult := Result{}
	err := c.sendRequest("POST", "add_result_for_case/"+strconv.Itoa(runID), newResult, &createdResult)
	return createdResult, err
}

func applyFiltersForCaseResults(uri string, filters RequestFilterForCaseResults) string {
	if filters.Limit != nil {
		uri = uri + "&limit=" + strconv.Itoa(*filters.Limit)
	}
	if filters.Offest != nil {
		uri = uri + "&offset=" + strconv.Itoa(*filters.Offest)
	}
	if len(filters.StatusID) != 0 {
		uri = applySpecificFilter(uri, "status_id", filters.StatusID)
	}

	return uri
}

func applyFiltersForRunResults(uri string, filters RequestFilterForRunResults) string {
	if filters.CreatedAfter != "" {
		uri = uri + "&created_after=" + filters.CreatedAfter
	}
	if filters.CreatedBefore != "" {
		uri = uri + "&created_before=" + filters.CreatedBefore
	}
	if len(filters.CreatedBy) != 0 {
		uri = applySpecificFilter(uri, "created_by", filters.CreatedBy)
	}
	if filters.Limit != nil {
		uri = uri + "&limit=" + strconv.Itoa(*filters.Limit)
	}
	if filters.Offest != nil {
		uri = uri + "&offset=" + strconv.Itoa(*filters.Offest)
	}
	if len(filters.StatusID) != 0 {
		uri = applySpecificFilter(uri, "status_id", filters.StatusID)
	}

	return uri
}
