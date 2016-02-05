package testrail

import (
	"fmt"
	"net/url"
)

// Milestone represents a Milestone
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

// SendableMilestone represents a Milestone
// that can be created or updated via the api
type SendableMilestone struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	DueOn       int    `json:"due_on,omitempty"`
}

// GetMilestone returns the existing milestone for a given milestoneID
func (c *Client) GetMilestone(milestoneID int) (ms Milestone, err error) {
	err = c.sendRequest("GET", fmt.Sprintf("get_milestone/%d", milestoneID), nil, &ms)
	return
}

// GetMilestones returns the list of milestones for the project projectID
// can be filtered by completed status of the milestones
func (c *Client) GetMilestones(projectID int, isCompleted ...bool) (milestones []Milestone, err error) {
	vals := make(url.Values)

	if len(isCompleted) > 0 {
		vals.Set("is_completed", boolToString(isCompleted[0]))
	}

	err = c.sendRequest("GET", fmt.Sprintf("get_milestones/%d?%s", projectID, vals.Encode()), nil, &milestones)
	return
}

// AddMilestone creates a new milestone on project projectID and returns it
func (c *Client) AddMilestone(projectID int, newMilestone SendableMilestone) (milestone Milestone, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("add_milestone/%d", projectID), newMilestone, &milestone)
	return
}

// UpdateMilestone updates the existing milestone milestoneID an returns it
func (c *Client) UpdateMilestone(milestoneID int, updates SendableMilestone) (milestone Milestone, err error) {
	err = c.sendRequest("POST", fmt.Sprintf("update_milestone/%d", milestoneID), updates, &milestone)
	return
}

// DeleteMilestone deletes the milestone milestoneID
func (c *Client) DeleteMilestone(milestoneID int) error {
	return c.sendRequest("POST", fmt.Sprintf("delete_milestone/%d", milestoneID), nil, nil)
}
