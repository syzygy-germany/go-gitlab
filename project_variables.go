package gitlab

import (
	"fmt"
	"net/url"
)

// ProjectVariablesService handles communication with the project variables
// related methods of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/project_level_variables.html
type ProjectVariablesService struct {
	client *Client
}

// ListProjectVariablesOptions represents the available ListSnippets() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/project_level_variables.html#list-project-variables
type ListProjectVariablesOptions ListOptions

// ListVariables gets a list of project variables.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/project_level_variables.html#list-project-variables
func (s *ProjectVariablesService) ListVariables(pid interface{}, opt *ListProjectVariablesOptions, options ...OptionFunc) ([]*Variable, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/variables", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var pv []*Variable
	resp, err := s.client.Do(req, &pv)
	if err != nil {
		return nil, resp, err
	}

	return pv, resp, err
}

// CreateProjectVariableOptions represents the available CreateVariable() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/project_level_variables.html#create-variable
type CreateProjectVariableOptions struct {
	Key       *string `url:"key,omitempty" json:"key,omitempty"`
	Value     *string `url:"value,omitempty" json:"value,omitempty"`
	Protected *bool   `url:"protected,omitempty" json:"protected,omitempty"`
}

// CreateVariable creates a new project snippet. The user must have permission
// to create new snippets.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/project_level_variables.html#create-variable
func (s *ProjectVariablesService) CreateVariable(pid interface{}, opt *CreateProjectVariableOptions, options ...OptionFunc) (*Variable, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/variables", url.QueryEscape(project))

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pv := new(Variable)
	resp, err := s.client.Do(req, pv)
	if err != nil {
		return nil, resp, err
	}

	return pv, resp, err
}
