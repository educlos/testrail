package testrail

import (
	"fmt"
	"net/url"
)

// Test represent a Test
type Test struct {
	AssignedToID     int    `json:"assignedto_id"`
	CaseID           int    `json:"case_id"`
	Estimate         string `json:"estimate"`
	EstimateForecast string `json:"estimate_forecast"`
	ID               int    `json:"id"`
	MilestoneID      int    `json:"milestone_id"`
	PriorityID       int    `json:"priority_id"`
	Refs             string `json:"refs"`
	RunID            int    `json:"run_id"`
	StatusID         int    `json:"status_id"`
	Title            string `json:"title"`
	TypeID           int    `json:"type_id"`
}

// GetTest returns the test testID
func (c *Client) GetTest(testID int) (test Test, err error) {
	err = c.sendRequest("GET", fmt.Sprintf("get_test/%d", testID), nil, &test)
	return
}

// GetTests returns the list of tests of runID
// with status statusID, if specified
func (c *Client) GetTests(runID int, statusID ...[]int) (tests []Test, err error) {
	vals := make(url.Values)

	if len(statusID) > 0 {
		vals.Set("status_id", intsList(statusID[0]))
	}

	err = c.sendRequest("GET", fmt.Sprintf("get_tests/%d?%s", runID, vals.Encode()), nil, &tests)
	return
}
