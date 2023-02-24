package prefect2_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetAllBlockDocuments(workspaceID string) ([]BlockDocument, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accounts/%s/workspaces/%s/block_documents/filter", c.PrefectApiUrl, c.PrefectAccountId, workspaceID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	blockDocuments := []BlockDocument{}

	err = json.Unmarshal(body, &blockDocuments)
	if err != nil {
		return nil, err
	}

	return blockDocuments, nil
}

func (c *Client) GetBlockDocument(blockDocumentId string, workspaceID string) (*BlockDocument, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/accounts/%s/workspaces/%s/block_documents/%s", c.PrefectApiUrl, c.PrefectAccountId, workspaceID, blockDocumentId), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	blockDocument := BlockDocument{}
	err = json.Unmarshal(body, &blockDocument)
	if err != nil {
		return nil, err
	}

	return &blockDocument, nil
}

func (c *Client) CreateBlockDocument(blockDocument BlockDocument, workspaceID string) (*BlockDocument, error) {
	rb, err := json.Marshal(blockDocument)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accounts/%s/workspaces/%s/block_documents/", c.PrefectApiUrl, c.PrefectAccountId, workspaceID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	}

	newBlockDocument := BlockDocument{}
	err = json.Unmarshal(body, &newBlockDocument)
	if err != nil {
		return nil, err
	}

	return &newBlockDocument, nil
}

func (c *Client) UpdateBlockDocument(blockDocument BlockDocument, blockDocumentId string, workspaceID string) (*BlockDocument, error) {
	rb, err := json.Marshal(blockDocument)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/accounts/%s/workspaces/%s/block_documents/%s", c.PrefectApiUrl, c.PrefectAccountId, workspaceID, blockDocumentId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, c.PrefectApiKey)
	if err != nil {
		return nil, err
	} else if len(body) == 0 {
		return nil, nil
	} else {
		updatedBlockDocument := BlockDocument{}
		err = json.Unmarshal(body, &updatedBlockDocument)
		if err != nil {
			return nil, err
		}

		return &updatedBlockDocument, nil
	}
}

func (c *Client) DeleteBlockDocument(blockDocumentId string, workspaceID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/accounts/%s/workspaces/%s/block_documents/%s", c.PrefectApiUrl, c.PrefectAccountId, workspaceID, blockDocumentId), nil)
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
