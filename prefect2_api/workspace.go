package prefect2_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetAllWorkspaces() ([]Workspace, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accounts/%s/workspaces/filter", c.PrefectApiUrl, c.PrefectAccountId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	workspaces := []Workspace{}

	err = json.Unmarshal(body, &workspaces)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}

func (c *Client) GetWorkspace(workspaceID string) (*Workspace, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/accounts/%s/workspaces/%s", c.PrefectApiUrl, c.PrefectAccountId, workspaceID), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	workspace := Workspace{}
	err = json.Unmarshal(body, &workspace)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (c *Client) CreateWorkspace(workspace Workspace) (*Workspace, error) {
	rb, err := json.Marshal(workspace)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accounts/%s/workspaces/", c.PrefectApiUrl, c.PrefectAccountId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	newWorkspace := Workspace{}
	err = json.Unmarshal(body, &newWorkspace)
	if err != nil {
		return nil, err
	}

	return &newWorkspace, nil
}

func (c *Client) DeleteWorkspace(workspaceID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/accounts/%s/workspaces/%s", c.PrefectApiUrl, c.PrefectAccountId, workspaceID), nil)

	if err != nil {
		return err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}
	
	return nil
}