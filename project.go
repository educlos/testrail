package testrail

import (
	"fmt"
	"strconv"
)

type Project struct {
	Announcement     string `json:"announcement"`
	CompletedOn      int    `json:"completed_on"`
	ID               int    `json:"id"`
	IsCompleted      bool   `json:"is_completed"`
	Name             string `json:"name"`
	ShowAnnouncement bool   `json:"show_announcement"`
	URL              string `json:"url"`
}

type SendableProject struct {
	Name             string `json:"name"`
	Announcement     string `json:"announcement,omitempty"`
	ShowAnnouncement bool   `json:"show_announcement,omitempty"`
	SuiteMode        int    `json:"suite_mode,omitempty"`
}

// Returns the existing project projectID
func (c *Client) GetProject(projectID int) (Project, error) {
	returnProject := Project{}
	err := c.sendRequest("GET", "get_project/"+strconv.Itoa(projectID), nil, &returnProject)
	return returnProject, err
}

// Returns a list available projects
func (c *Client) GetProjects(isCompleted ...bool) ([]Project, error) {
	uri := "get_projects"
	if len(isCompleted) > 0 {
		uri = uri + "&is_completed=" + btoitos(isCompleted[0])
	}

	returnProjects := []Project{}
	err := c.sendRequest("GET", uri, nil, &returnProjects)
	return returnProjects, err
}

// Creates a new project and return the created project
func (c *Client) AddProject(newProject SendableProject) (Project, error) {
	createdProject := Project{}
	err := c.sendRequest("POST", "add_project", newProject, &createdProject)
	return createdProject, err
}

// Updates the existing project projectID
func (c *Client) UpdateProject(projectID int, updates SendableProject, isCompleted ...bool) (Project, error) {
	updatedProject := Project{}
	uri := "update_project/" + strconv.Itoa(projectID)
	if len(isCompleted) > 0 {
		uri = uri + "&is_completed=" + btoitos(isCompleted[0])
	}

	fmt.Println(uri)
	err := c.sendRequest("POST", uri, updates, &updatedProject)
	return updatedProject, err
}

// Deletes the existing project projectID
func (c *Client) DeleteProject(projectID int) error {
	return c.sendRequest("POST", "delete_project/"+strconv.Itoa(projectID), nil, nil)
}
