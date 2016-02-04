package testrail

import (
	"fmt"
	"net/url"
)

// Project represents a Project
type Project struct {
	Announcement     string `json:"announcement"`
	CompletedOn      int    `json:"completed_on"`
	ID               int    `json:"id"`
	IsCompleted      bool   `json:"is_completed"`
	Name             string `json:"name"`
	ShowAnnouncement bool   `json:"show_announcement"`
	URL              string `json:"url"`
}

// SendableProject represents a Project
// that can be created or updated via the ap
type SendableProject struct {
	Name             string `json:"name"`
	Announcement     string `json:"announcement,omitempty"`
	ShowAnnouncement bool   `json:"show_announcement,omitempty"`
	SuiteMode        int    `json:"suite_mode,omitempty"`
}

// GetProject returns the existing project projectID
func (c *Client) GetProject(projectID int) (project Project, err error) {
	err = c.sendRequest("GET", fmt.Sprintf("get_project/%d", projectID), nil, &project)
	return
}

// GetProjects returns a list available projects
// can be filtered by completed status of the project
func (c *Client) GetProjects(isCompleted ...bool) (projects []Project, err error) {
	vals := make(url.Values)

	if len(isCompleted) > 0 {
		vals.Set("is_completed", boolToString(isCompleted[0]))
	}

	err = c.sendRequest("GET", fmt.Sprintf("get_projects?%s", vals.Encode()), nil, &projects)
	return
}

// AddProject creates a new project and return its
func (c *Client) AddProject(newProject SendableProject) (project Project, err error) {
	err = c.sendRequest("POST", "add_project", newProject, &project)
	return
}

// UpdateProject updates the existing project projectID and returns it
func (c *Client) UpdateProject(projectID int, updates SendableProject, isCompleted ...bool) (project Project, err error) {
	vals := make(url.Values)
	if len(isCompleted) > 0 {
		vals.Set("is_completed", boolToString(isCompleted[0]))
	}

	err = c.sendRequest("POST", fmt.Sprintf("update_project/%d?%s", projectID, vals.Encode()), updates, &project)
	return
}

// DeleteProject deletes the project projectID
func (c *Client) DeleteProject(projectID int) error {
	return c.sendRequest("POST", fmt.Sprintf("delete_project/%d", projectID), nil, nil)
}
