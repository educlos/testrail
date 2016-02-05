package testrail

import (
	"fmt"
	"net/url"
)

// Result represents a Test Case result
type Result struct {
	AssignedToID      int                `json:"assignedto_id"`
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

// CustomStepResult represents the custom steps
// results a Result can have
type CustomStepResult struct {
	Content  string `json:"content"`
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
	StatusID int    `json:"status_id"`
}

// RequestFilterForCaseResults represents the filters
// usable to get the test case results
type RequestFilterForCaseResults struct {
	Limit    *int    `json:"limit,omitempty"`
	Offset   *int    `json:"offset,omitempty"`
	StatusID IntList `json:"status_id,omitempty"`
}

// RequestFilterForRunResults represents the filters
// usable to get the run results
type RequestFilterForRunResults struct {
	CreatedAfter  string  `json:"created_after,omitempty"`
	CreatedBefore string  `json:"created_before,omitempty"`
	CreatedBy     IntList `json:"created_by,omitempty"`
	Limit         *int    `json:"limit,omitempty"`
	Offset        *int    `json:"offset,omitempty"`
	StatusID      IntList `json:"status_id,omitempty"`
}

// SendableResult represents a Test Case result
// that can be created or updated via the api
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
	MultiSelect  []int              `json:"custom_multi_select,omitempty"`
	StepsResults []CustomStepResult `json:"custom_step_results,omitempty"`
	String       string             `json:"custom_string,omitempty"`
	Text         string             `json:"custom_text,omitempty"`
	URL          string             `json:"custom_url,omitempty"`
	User         int                `json:"custom_user,omitempty"`
}

// SendableResults represents a list of run results
// that can be created or updated via the api
type SendableResults struct {
	Results []Results `json:"results"`
}

// Results represents a run result
// that can be created or updated via the api
type Results struct {
	TestID int `json:"test_id"`
	SendableResult
}

// SendableResultsForCase represents a Test Case result
// that can be created or updated via the api
type SendableResultsForCase struct {
	Results []Results `json:"results"`
}

// GetResults returns a list of results for the test testID
// validating the filters
func (c *Client) GetResults(testID int, filters ...RequestFilterForCaseResults) (results []Result, err error) {
	vals := make(url.Values)
	loadOptionalFilters(vals, filters)

	err = c.sendRequest("GET", fmt.Sprintf("get_results/%d?%s", testID, vals.Encode()), nil, &results)
	return
}

// GetResultsForCase returns a list of results for the case caseID
// on run runID validating the filters
func (c *Client) GetResultsForCase(runID, caseID int, filters ...RequestFilterForCaseResults) (results []Result, err error) {
	vals := make(url.Values)
	loadOptionalFilters(vals, filters)

	err = c.sendRequest("GET", fmt.Sprintf("get_results_for_case/%d/%d?%s", runID, caseID, vals.Encode()), nil, &results)
	return
}

// GetResultsForRun returns a list of results for the run runID
// validating the filters
func (c *Client) GetResultsForRun(runID int, filters ...RequestFilterForRunResults) (results []Result, err error) {
	vals := make(url.Values)
	loadOptionalFilters(vals, filters)

	err = c.sendRequest("GET", fmt.Sprintf("get_results_for_run/%d?%s", runID, vals.Encode()), nil, &results)
	return
}

// AddResult adds a new result, comment or assigns a test to testID
func (c *Client) AddResult(testID int, newResult SendableResult) (result Result, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("add_result/%d", testID), newResult, &result)
	return
}

// AddResultForCase adds a new result, comment or assigns a test to the case caseID on run runID
func (c *Client) AddResultForCase(runID, caseID int, newResult SendableResult) (result Result, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("add_result_for_case/%d/%d", runID, caseID), newResult, &result)
	return
}

// AddResults adds new results, comment or assigns tests to runID
func (c *Client) AddResults(runID int, newResult SendableResults) (result []Result, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("add_results/%d", runID), newResult, &result)
	return
}

// AddResultsForCase adds new results, comments or assigns tests to run runID
// each result being assigned to a test case
func (c *Client) AddResultsForCase(runID int, newResult SendableResultsForCase) (result Result, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("add_result_for_case/%d", runID), newResult, &result)
	return
}
