package testrail

import "strconv"

type Plan struct {
	AssignedToID  int     `json:"assignedto_id"`
	BlockedCount  int     `json:"blocked_count"`
	CompletedOn   int     `json:"completed_on"`
	CreatedBy     int     `json:"created_by"`
	CreatedOn     int     `json:"created_on"`
	Description   string  `json:"description"`
	Entries       []Entry `json:"entries"`
	FailedCount   int     `json:"failed_count"`
	ID            int     `json:"id"`
	IsCompleted   bool    `json:"is_completed"`
	MilestoneID   int     `json:"milestone_id"`
	Name          string  `json:"name"`
	PassedCount   int     `json:"passed_count"`
	ProjectID     int     `json:"project_id"`
	RetestCount   int     `json:"retest_count"`
	UntestedCount int     `json:"untested_count"`
	URL           string  `json:"url"`
}

type Entry struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Runs    []Run  `json:"runs"`
	SuiteID int    `json:"suite_id"`
}

type Run struct {
	AssignedToID  int    `json:"assignedto_id"`
	BlockedCount  int    `json:"blocked_count"`
	CompletedOn   int    `json:"completed_on"`
	Config        string `json:"config"`
	ConfigIDs     []int  `json:"config_ids"`
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

type RequestFilterForPlan struct {
	CreatedAfter  string `json:"created_after"`
	CreatedBefore string `json:"created_before"`
	CreatedBy     []int  `json:"created_by"`
	IsCompleted   *bool  `json:"is_completed"`
	Limit         *int   `json:"limit"`
	Offest        *int   `json:"offset"`
	MilestoneId   []int  `json:"milestone_id"`
}

type SendablePlan struct {
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	MilestoneID int               `json:"milestone_id,omitempty"`
	Entries     []SendableEntries `json:"entries,omitempty"`
}

type SendableEntries struct {
	SuiteID      int    `json:"suite_id"`
	Name         string `json:"name,omitempty"`
	AssignedtoID int    `json:"assignedto_id,omitempty"`
	IncludeAll   bool   `json:"include_all,omitempty"`
	CaseIds      []int  `json:"case_ids,omitempty"`
	ConfigIds    []int  `json:"config_ids,omitempty"`
	Runs         []Run  `json:"runs,omitempty"`
}

// Returns the existing plan planID
func (c *Client) GetPlan(planID int) (Plan, error) {
	returnPlan := Plan{}
	err := c.sendRequest("GET", "get_plan/"+strconv.Itoa(planID), nil, &returnPlan)
	return returnPlan, err
}

// Returns the list of plans for a project.
func (c *Client) GetPlans(projectID int, filters ...RequestFilterForPlan) ([]Plan, error) {
	uri := "get_plans/" + strconv.Itoa(projectID)
	if len(filters) > 0 {
		uri = applyFiltersForPlan(uri, filters[0])
	}

	returnPlans := []Plan{}
	err := c.sendRequest("GET", uri, nil, &returnPlans)
	return returnPlans, err
}

// Creates a new plan and return the created plan
func (c *Client) AddPlan(projectID int, newPlan SendablePlan) (Plan, error) {
	createdPlan := Plan{}
	err := c.sendRequest("POST", "add_plan/"+strconv.Itoa(projectID), newPlan, &createdPlan)
	return createdPlan, err
}

// Creates a new entry on plan planID and return the created entry
func (c *Client) AddPlanEntry(planID int, newEntry SendableEntries) (Entry, error) {
	createdEntry := Entry{}
	err := c.sendRequest("POST", "add_plan_entry/"+strconv.Itoa(planID), newEntry, &createdEntry)
	return createdEntry, err
}

// Updates the existing plan planID
func (c *Client) UpdatePlan(planID int, updates SendablePlan) (Plan, error) {
	updatedPlan := Plan{}
	err := c.sendRequest("POST", "update_plan/"+strconv.Itoa(planID), updates, &updatedPlan)
	return updatedPlan, err
}

// Update the entry entryID on plan planID
func (c *Client) UpdatePlanEntry(planID int, entryID string, updates SendableEntries) (Entry, error) {
	uri := "update_plan_entry/" + strconv.Itoa(planID) + "/" + entryID
	updatedEntry := Entry{}
	err := c.sendRequest("POST", uri, updates, &updatedEntry)
	return updatedEntry, err
}

// Close the plan planId
func (c *Client) ClosePlan(planID int) (Plan, error) {
	deletedPlan := Plan{}
	err := c.sendRequest("POST", "close_plan/"+strconv.Itoa(planID), nil, &deletedPlan)
	return deletedPlan, err
}

// Deletes the existing plan planID
func (c *Client) DeletePlan(planID int) error {
	return c.sendRequest("POST", "delete_plan/"+strconv.Itoa(planID), nil, nil)
}

// Delete the entry entryID on plan planID
func (c *Client) DeletePlanEntry(planID int, entryID string) error {
	uri := "delete_plan_entry/" + strconv.Itoa(planID) + "/" + entryID
	return c.sendRequest("POST", uri, nil, nil)
}

func applyFiltersForPlan(uri string, filters RequestFilterForPlan) string {
	if filters.CreatedAfter != "" {
		uri = uri + "&created_after=" + filters.CreatedAfter
	}
	if filters.CreatedBefore != "" {
		uri = uri + "&created_before=" + filters.CreatedBefore
	}
	if len(filters.CreatedBy) != 0 {
		uri = applySpecificFilter(uri, "created_by", filters.CreatedBy)
	}
	if len(filters.MilestoneId) != 0 {
		uri = applySpecificFilter(uri, "milestone_id", filters.MilestoneId)
	}
	if filters.IsCompleted != nil {
		uri = uri + "&is_completed=" + btoitos(*filters.IsCompleted)
	}
	if filters.Limit != nil {
		uri = uri + "&limit=" + strconv.Itoa(*filters.Limit)
	}
	if filters.Offest != nil {
		uri = uri + "&offset=" + strconv.Itoa(*filters.Offest)
	}

	return uri
}
