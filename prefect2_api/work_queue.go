package prefect2_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetAllWorkQueues(workspaceID string) ([]WorkQueue, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accounts/%s/workspaces/%s/work_queues/filter", c.PrefectApiUrl, c.PrefectAccountId, workspaceID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	workQueues := []WorkQueue{}

	err = json.Unmarshal(body, &workQueues)
	if err != nil {
		return nil, err
	}

	return workQueues, nil
}

func (c *Client) GetWorkQueue(workQueueId string, workspaceID string) (*WorkQueue, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/accounts/%s/workspaces/%s/work_queues/%s", c.PrefectApiUrl, c.PrefectAccountId, workspaceID, workQueueId), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	workQueue := WorkQueue{}
	err = json.Unmarshal(body, &workQueue)
	if err != nil {
		return nil, err
	}

	return &workQueue, nil
}

func (c *Client) CreateWorkQueue(workQueue WorkQueue, workspaceID string) (*WorkQueue, error) {
	rb, err := json.Marshal(workQueue)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accounts/%s/workspaces/%s/work_queues/", c.PrefectApiUrl, c.PrefectAccountId, workspaceID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	newWorkQueue := WorkQueue{}
	err = json.Unmarshal(body, &newWorkQueue)
	if err != nil {
		return nil, err
	}

	return &newWorkQueue, nil
}

func (c *Client) UpdateWorkQueue(workQueue WorkQueue, workQueueId string, workspaceID string) (*WorkQueue, error) {
	rb, err := json.Marshal(workQueue)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/accounts/%s/workspaces/%s/work_queues/%s", c.PrefectApiUrl, c.PrefectAccountId, workspaceID, workQueueId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	} else if len(body) == 0 {
		return nil, nil
	} else {
		updatedWorkQueue := WorkQueue{}
		err = json.Unmarshal(body, &updatedWorkQueue)
		if err != nil {
			return nil, err
		}

		return &updatedWorkQueue, nil
	}
}

func (c *Client) DeleteWorkQueue(workQueueId string, workspaceID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/accounts/%s/workspaces/%s/work_queues/%s", c.PrefectApiUrl, c.PrefectAccountId, workspaceID, workQueueId), nil)
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
