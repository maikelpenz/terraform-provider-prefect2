package prefect2_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetAllBlockTypes(workspaceID string) ([]BlockType, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accounts/%s/workspaces/%s/block_types/filter", c.PrefectApiUrl, c.PrefectAccountId, workspaceID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	blockTypes := []BlockType{}

	err = json.Unmarshal(body, &blockTypes)
	if err != nil {
		return nil, err
	}

	return blockTypes, nil
}

func (c *Client) GetBlockTypeById(blockTypeId string, workspaceID string) (*BlockType, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/accounts/%s/workspaces/%s/block_types/%s", c.PrefectApiUrl, c.PrefectAccountId, workspaceID, blockTypeId), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	blockType := BlockType{}
	err = json.Unmarshal(body, &blockType)
	if err != nil {
		return nil, err
	}

	return &blockType, nil
}

func (c *Client) GetBlockTypeBySlug(slug string, workspaceID string) (*BlockType, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/accounts/%s/workspaces/%s/block_types/slug/%s", c.PrefectApiUrl, c.PrefectAccountId, workspaceID, slug), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	blockType := BlockType{}
	err = json.Unmarshal(body, &blockType)
	if err != nil {
		return nil, err
	}

	return &blockType, nil
}
