package testrail

import "strconv"

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

// Returns the existing test testID
func (c *Client) GetTest(testID int) (Test, error) {
	returnTest := Test{}
	err := c.sendRequest("GET", "get_test/"+strconv.Itoa(testID), nil, &returnTest)
	return returnTest, err
}

// Returns the list of tests of runID, with status statusID, if specified
func (c *Client) GetTests(runID int, statusID ...[]int) ([]Test, error) {
	returnTest := []Test{}
	uri := "get_tests/" + strconv.Itoa(runID)

	if len(statusID) > 0 {
		uri = applySpecificFilter(uri, "status_id", statusID[0])
	}
	err := c.sendRequest("GET", uri, nil, &returnTest)
	return returnTest, err
}
