package testrail

import "strconv"

type Milestone struct {
	CompletedOn int    `json:"completed_on"`
	Description string `json:"description"`
	DueOn       int    `json:"due_on"`
	ID          int    `json:"id"`
	IsCompleted bool   `json:"is_completed"`
	Name        string `json:"name"`
	ProjectID   int    `json:"project_id"`
	URL         string `json:"url"`
}

type SendableMilestone struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	DueOn       int    `json:"due_on,omitempty"`
}

// Returns the existing milestone milestoneID
func (c *Client) GetMilestone(milestoneID int) (Milestone, error) {
	returnMilestone := Milestone{}
	err := c.sendRequest("GET", "get_milestone/"+strconv.Itoa(milestoneID), nil, &returnMilestone)
	return returnMilestone, err
}

// Returns the list of milestones for a project.
func (c *Client) GetMilestones(projectID int, isCompleted ...bool) ([]Milestone, error) {
	uri := "get_milestones/" + strconv.Itoa(projectID)
	if len(isCompleted) > 0 {
		uri = uri + "&is_completed=" + btoitos(isCompleted[0])
	}

	returnMilestones := []Milestone{}
	err := c.sendRequest("GET", uri, nil, &returnMilestones)
	return returnMilestones, err
}

// Creates a new milestone and return the created milestone
func (c *Client) AddMilestone(projectID int, newMilestone SendableMilestone) (Milestone, error) {
	createdMilestone := Milestone{}
	err := c.sendRequest("POST", "add_milestone/"+strconv.Itoa(projectID), newMilestone, &createdMilestone)
	return createdMilestone, err
}

// Updates the existing milestone milestoneID
func (c *Client) UpdateMilestone(milestoneID int, updates SendableMilestone) (Milestone, error) {
	updatedMilestone := Milestone{}
	err := c.sendRequest("POST", "update_milestone/"+strconv.Itoa(milestoneID), updates, &updatedMilestone)
	return updatedMilestone, err
}

// Deletes the existing milestone milestoneID
func (c *Client) DeleteMilestone(milestoneID int) error {
	return c.sendRequest("POST", "delete_milestone/"+strconv.Itoa(milestoneID), nil, nil)
}
