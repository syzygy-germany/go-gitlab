package gitlab

// Variable represents a GitLab project or group variable. Note: when
// retrieving variables (details or not), Protected is
//
// GitLab API docs: https://docs.gitlab.com/ce/api/project_level_variables.html
// GitLab API docs: https://docs.gitlab.com/ee/api/group_level_variables.html
type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
