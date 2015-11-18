package testrail

type CaseField struct {
	Configs      []CaseFieldConfig `json:"configs"`
	Description  string            `json:"description"`
	DisplayOrder int               `json:"display_order"`
	ID           int               `json:"ID"`
	IsActive     bool              `json:"is_active"`
	Label        string            `json:"label"`
	Name         string            `json:"name"`
	SystemName   string            `json:"system_name"`
	TypeID       int               `json:"type_id"`
}

type CaseFieldConfig struct {
	Context Context `json:"context"`
	ID      string  `json:"id"`
	Options Option  `json:"options"`
}

type Context struct {
	IsGlobal   bool  `json:"is_global"`
	ProjectIDs []int `json:"project_ids"`
}

type Option struct {
	DefaultValue string `json:"default_value"`
	Format       string `json:"format"`
	IsRequired   bool   `json:"is_required"`
	Rows         string `json:"rows"`
}

// Returns a list of available case types.
func (c *Client) GetCaseFields() ([]CaseField, error) {
	caseFields := []CaseField{}
	err := c.sendRequest("GET", "get_case_fields", nil, &caseFields)
	return caseFields, err
}
