package testrail

import "fmt"

type Configuration struct {
	Configs   []Config `json:"configs"`
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	ProjectID int      `json:"project_id"`
}

type Config struct {
	GroupID int    `json:"group_id"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
}

// GetConfigs returns a list of available configurations, grouped by configuration groups
func (c *Client) GetConfigs(projectID int) (configs []Configuration, err error) {
	err = c.sendRequest("GET", fmt.Sprintf("get_configs/%d", projectID), nil, &configs)
	return
}
