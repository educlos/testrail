package testrail

import "strconv"

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

// Returns a list of available configurations, grouped by configuration groups
func (c *Client) GetCongifs(projectID int) ([]Configuration, error) {
	configs := []Configuration{}
	err := c.sendRequest("GET", "get_configs/"+strconv.Itoa(projectID), nil, &configs)
	return configs, err
}
