package testrail

import (
	"fmt"
	"net/url"
)

// Plan represents a Plan
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

// Entry represents the entry a Plan can have
type Entry struct {
	ID      string `json:"id"` // these are GUIDs, unlike the other IDs around.
	Name    string `json:"name"`
	Runs    []Run  `json:"runs"`
	SuiteID int    `json:"suite_id"`
}

// RequestFilterForPlan represents the filters
// usable to get the plan
type RequestFilterForPlan struct {
	CreatedAfter  string  `json:"created_after,omitempty"`
	CreatedBefore string  `json:"created_before,omitempty"`
	CreatedBy     IntList `json:"created_by,omitempty"`
	IsCompleted   *bool   `json:"is_completed,omitempty"`
	Limit         *int    `json:"limit,omitempty"`
	Offset        *int    `json:"offset,omitempty"`
	MilestoneID   IntList `json:"milestone_id,omitempty"`
}

// SendablePlan represents a Plan
// that can be created or updated via the api
type SendablePlan struct {
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	MilestoneID int             `json:"milestone_id,omitempty"`
	Entries     []SendableEntry `json:"entries,omitempty"`
}

// SendableEntry represents an Entry
// that can be created or updated via the api
type SendableEntry struct {
	SuiteID      int    `json:"suite_id"`
	Name         string `json:"name,omitempty"`
	AssignedToID int    `json:"assignedto_id,omitempty"`
	IncludeAll   bool   `json:"include_all,omitempty"`
	CaseIDs      []int  `json:"case_ids,omitempty"`
	ConfigIDs    []int  `json:"config_ids,omitempty"`
	Runs         []Run  `json:"runs,omitempty"`
}

// GetPlan returns the existing plan planID
func (c *Client) GetPlan(planID int) (plan Plan, err error) {
	err = c.sendRequest("GET", fmt.Sprintf("get_plan/%d", planID), nil, &plan)
	return
}

// GetPlans returns the list of plans for the project projectID
// validating the filters
func (c *Client) GetPlans(projectID int, filters ...RequestFilterForPlan) (plans []Plan, err error) {
	vals := make(url.Values)

	loadOptionalFilters(vals, filters)

	err = c.sendRequest("GET", fmt.Sprintf("get_plans/%d?%s", projectID, vals.Encode()), nil, &plans)
	return
}

// AddPlan creates a new plan on project projectID and returns it
func (c *Client) AddPlan(projectID int, newPlan SendablePlan) (plan Plan, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("add_plan/%d", projectID), newPlan, &plan)
	return
}

// AddPlanEntry creates a new entry on plan planID and returns it
func (c *Client) AddPlanEntry(planID int, newEntry SendableEntry) (entry Entry, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("add_plan_entry/%d", planID), newEntry, &entry)
	return
}

// UpdatePlan updates the existing plan planID and returns it
func (c *Client) UpdatePlan(planID int, updates SendablePlan) (plan Plan, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("update_plan/%d", planID), updates, &plan)
	return
}

// UpdatePlanEntry updates the entry entryID on plan planID and returns it
func (c *Client) UpdatePlanEntry(planID int, entryID string, updates SendableEntry) (entry Entry, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("update_plan_entry/%d/%s", planID, entryID), updates, &entry)
	return
}

// ClosePlan closes the plan planID and returns it
func (c *Client) ClosePlan(planID int) (plan Plan, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("close_plan/%d", planID), nil, &plan)
	return
}

// DeletePlan deletes the plan planID
func (c *Client) DeletePlan(planID int) error {
	return c.sendRequest("POST", fmt.Sprintf("delete_plan/%d", planID), nil, nil)
}

// DeletePlanEntry delete the entry entryID on plan planID
func (c *Client) DeletePlanEntry(planID int, entryID string) error {
	return c.sendRequest("POST", fmt.Sprintf("delete_plan_entry/%d/%s", planID, entryID), nil, nil)
}
