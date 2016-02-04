package testrail

import (
	"fmt"
	"net/url"
)

// Run represents a Run
type Run struct {
	AssignedToID  int    `json:"assignedto_id"`
	BlockedCount  int    `json:"blocked_count"`
	CompletedOn   int    `json:"completed_on"`
	Config        string `json:"config"`
	ConfigIDs     []int  `json:"config_ids"`
	CreatedBy     int    `json:"created_by"`
	CreatedOn     int    `json:"created_on"`
	Description   string `json:"description"`
	EntryID       string `json:"entry_id"`
	EntryIndex    int    `json:"entry_index"`
	FailedCount   int    `json:"failed_count"`
	ID            int    `json:"id"`
	IncludeAll    bool   `json:"include_all"`
	IsCompleted   bool   `json:"is_completed"`
	MilestoneID   int    `json:"milestone_id"`
	Name          string `json:"name"`
	PassedCount   int    `json:"passed_count"`
	PlanID        int    `json:"plan_id"`
	ProjectID     int    `json:"project_id"`
	RetestCount   int    `json:"retest_count"`
	SuiteID       int    `json:"suite_id"`
	UntestedCount int    `json:"untested_count"`
	URL           string `json:"url"`
}

// RequestFilterForRun represents the filters
// usable to get the run
type RequestFilterForRun struct {
	CreatedAfter  string  `json:"created_after,omitempty"`
	CreatedBefore string  `json:"created_before,omitempty"`
	CreatedBy     IntList `json:"created_by,omitempty"`
	IsCompleted   *bool   `json:"is_completed,omitempty"`
	Limit         *int    `json:"limit,omitempty"`
	Offset        *int    `json:"offset,omitempty"`
	MilestoneID   IntList `json:"milestone_id,omitempty"`
	SuiteID       IntList `json:"suite_id,omitempty"`
}

// SendableRun represents a Run
// that can be created via the api
type SendableRun struct {
	SuiteID      int    `json:"suite_id"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	MilestoneID  int    `json:"milestone_id,omitempty"`
	AssignedToID int    `json:"assignedto_id,omitempty"`
	IncludeAll   *bool  `json:"include_all,omitempty"`
	CaseIDs      []int  `json:"case_id,omitempty"`
}

// UpdatableRun represents a Run
// that can be updated via the api
type UpdatableRun struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	MilestoneID int    `json:"milestone_id,omitempty"`
	IncludeAll  *bool  `json:"include_all,omitempty"`
	CaseIDs     []int  `json:"case_id,omitempty"`
}

// GetRun returns the run runID
func (c *Client) GetRun(runID int) (run Run, err error) {
	err = c.sendRequest("GET", fmt.Sprintf("get_run/%d", runID), nil, &run)
	return
}

// GetRuns returns the list of runs of projectID
// validating the filters
func (c *Client) GetRuns(projectID int, filters ...RequestFilterForRun) (runs []Run, err error) {
	vals := make(url.Values)
	loadOptionalFilters(vals, filters)
	err = c.sendRequest("GET", fmt.Sprintf("get_runs/%d?%s", projectID, vals.Encode()), nil, &runs)
	return
}

// AddRun creates a new run on projectID and returns it
func (c *Client) AddRun(projectID int, newRun SendableRun) (run Run, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("add_run/%d", projectID), newRun, &run)
	return
}

// UpdateRun updates the run runID and returns it
func (c *Client) UpdateRun(runID int, update UpdatableRun) (run Run, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("update_run/%d", runID), update, &run)
	return
}

// CloseRun closes the run runID,
// archives its tests and results
// and returns it
func (c *Client) CloseRun(runID int) (run Run, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("close_run/%d", runID), nil, &run)
	return
}

// DeleteRun delete the run runID
func (c *Client) DeleteRun(runID int) error {
	return c.sendRequest("POST", fmt.Sprintf("delete_run/%d", runID), nil, nil)
}
