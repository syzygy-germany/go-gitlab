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

// ListSnippets gets a list of project variables.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/project_level_variables.html#list-project-variables
func (s *ProjectVariablesService) ListSnippets(pid interface{}, opt *ListProjectVariablesOptions, options ...OptionFunc) ([]*Variable, *Response, error) {
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
