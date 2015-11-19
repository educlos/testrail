package testrail

import "strconv"

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

type RequestFilterForRun struct {
	CreatedAfter  string `json:"created_after,omitempty"`
	CreatedBefore string `json:"created_before,omitempty"`
	CreatedBy     []int  `json:"created_by,omitempty"`
	IsCompleted   *bool  `json:"is_completed,omitempty"`
	Limit         *int   `json:"limit,omitempty"`
	Offset        *int   `json:"offset, omitempty"`
	MilestoneID   []int  `json:"milestone_id,omitempty"`
	SuiteID       []int  `json:"suite_id,omitempty"`
}

type SendableRun struct {
	SuiteID      int    `json:"suite_id"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	MilestoneID  int    `json:"milestone_id,omitempty"`
	AssignedToID int    `json:"assignedto_id,omitempty"`
	IncludeAll   *bool  `json:"include_all,omitempty"`
	CaseIDs      []int  `json:"case_id,omitempty"`
}

type UpdatableRun struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	MilestoneID int    `json:"milestone_id,omitempty"`
	IncludeAll  *bool  `json:"include_all,omitempty"`
	CaseIDs     []int  `json:"case_id,omitempty"`
}

// Returns the existing test run runID
func (c *Client) GetRun(runID int) (Run, error) {
	returnRun := Run{}
	err := c.sendRequest("GET", "get_run/"+strconv.Itoa(runID), nil, &returnRun)
	return returnRun, err
}

// Returns the list of test runs of projectID
func (c *Client) GetRuns(projectID int, filters ...RequestFilterForRun) ([]Run, error) {
	returnRun := []Run{}
	err := c.sendRequest("GET", "get_runs/"+strconv.Itoa(projectID), nil, &returnRun)
	return returnRun, err
}

// Creates a new test run on projectID and return the created test run
func (c *Client) AddRun(projectID int, newRun SendableRun) (Run, error) {
	createdRun := Run{}
	err := c.sendRequest("POST", "add_run/"+strconv.Itoa(projectID), newRun, &createdRun)
	return createdRun, err
}

// Updates the existing test run runID
func (c *Client) UpdateRun(runID int, update UpdatableRun) (Run, error) {
	updatedRun := Run{}
	err := c.sendRequest("POST", "update_run/"+strconv.Itoa(runID), update, &updatedRun)
	return updatedRun, err
}

// Close the existing test run runID and archives its tests & results
func (c *Client) CloseRun(runID int) (Run, error) {
	closedRun := Run{}
	err := c.sendRequest("POST", "close_run/"+strconv.Itoa(runID), nil, &closedRun)
	return closedRun, err
}

// Delete the existing test run runID
func (c *Client) DeleteRun(runID int) error {
	return c.sendRequest("POST", "delete_run/"+strconv.Itoa(runID), nil, nil)
}

func applyFiltersForRuns(uri string, filters RequestFilterForRun) string {
	if filters.CreatedAfter != "" {
		uri = uri + "&created_after=" + filters.CreatedAfter
	}
	if filters.CreatedBefore != "" {
		uri = uri + "&created_before=" + filters.CreatedBefore
	}
	if len(filters.CreatedBy) != 0 {
		uri = applySpecificFilter(uri, "created_by", filters.CreatedBy)
	}
	if filters.IsCompleted != nil {
		uri = uri + "&is_completed=" + btoitos(*filters.IsCompleted)
	}
	if filters.Limit != nil {
		uri = uri + "&limit=" + strconv.Itoa(*filters.Limit)
	}
	if filters.Offset != nil {
		uri = uri + "&offset=" + strconv.Itoa(*filters.Offset)
	}
	if len(filters.MilestoneID) != 0 {
		uri = applySpecificFilter(uri, "milestone_id", filters.MilestoneID)
	}
	if len(filters.SuiteID) != 0 {
		uri = applySpecificFilter(uri, "suite_id", filters.SuiteID)
	}

	return uri
}
