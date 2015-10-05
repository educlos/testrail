package TestRailGo

import "strconv"

type Case struct {
	CreatedBy      string `json:"created_by"`
	CreatedOn      int    `json:"created_on"`
	CustomExpected string `json:"custom_expected"`
	CustomPreconds string `json:"custom_preconds"`
	CustomSteps    string `json:"custom_steps"`
	Estimate       string `json:"estimate"`
	ID             int    `json"id"`
	MilestoneId    int    `json:"milestone_id"`
	PriorityId     int    `json:"priority_id"`
	Refs           string `json:"refs"`
	SectionId      int    `json:"section_id"`
	SuiteId        int    `json:"suite_id"`
	Title          string `json:"title"`
	TypeId         int    `json:"type_id"`
	UpdatedBy      int    `json:"updated_by"`
	UdpatedOn      int    `json:"updated_on"`
}

func (c *Client) GetCase(ID int) (Case, error) {
	return c.sendRequest("GET", "get_case/"+strconv.Itoa(ID), nil)
}

func (c *Client) GetCases(project, suite, section int) ([]Case, error) {
	uri := "get_cases/" + strconv.Itoa(project) + "&suite_id=" + strconv.Itoa(suite) + "&section_id=" + strconv.Itoa(section)
	return c.sendRequest("GET", uri, nil)
}
