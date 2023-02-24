package prefect2_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetAllBlockSchemas(workspaceID string) ([]BlockSchema, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accounts/%s/workspaces/%s/block_schemas/filter", c.PrefectApiUrl, c.PrefectAccountId, workspaceID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	blockSchemas := []BlockSchema{}

	err = json.Unmarshal(body, &blockSchemas)
	if err != nil {
		return nil, err
	}

	return blockSchemas, nil
}

func (c *Client) GetBlockSchemaById(blockSchemaId string, workspaceID string) (*BlockSchema, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/accounts/%s/workspaces/%s/block_schemas/%s", c.PrefectApiUrl, c.PrefectAccountId, workspaceID, blockSchemaId), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	blockSchema := BlockSchema{}
	err = json.Unmarshal(body, &blockSchema)
	if err != nil {
		return nil, err
	}

	return &blockSchema, nil
}

func (c *Client) GetBlockSchemaByChecksum(checksum string, workspaceID string) (*BlockSchema, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/accounts/%s/workspaces/%s/block_schemas/checksum/%s", c.PrefectApiUrl, c.PrefectAccountId, workspaceID, checksum), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	blockSchema := BlockSchema{}
	err = json.Unmarshal(body, &blockSchema)
	if err != nil {
		return nil, err
	}

	return &blockSchema, nil
}
