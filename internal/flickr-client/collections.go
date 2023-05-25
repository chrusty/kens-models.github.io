package flickrclient

import (
	"encoding/json"
	"fmt"

	"github.com/mncaudill/go-flickr"
)

type CollectionSet struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Collection struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	IconSmall   string          `json:"iconsmall"`
	Set         []CollectionSet `json:"set"`
}

type Collections struct {
	Collection []Collection `json:"collection"`
}

type CollectionsGetTreeResponse struct {
	Collections Collections `json:"collections"`
	Status      string      `json:"stat"`
}

func (c *Client) CollectionsGetTree() (*CollectionsGetTreeResponse, error) {

	// Build the request:
	req := &flickr.Request{
		ApiKey: c.apiKey,
		Method: methodCollectionsGetTree,
		Args: map[string]string{
			argUserId:         c.userId,
			argFormat:         c.format,
			argNoJSONCallBack: "1",
		},
	}

	// Sign the request:
	if len(c.secret) > 0 {
		req.Sign(c.secret)
	}

	// Submit the request:
	rspJSON, err := req.Execute()
	if err != nil {
		return nil, err
	}

	// Unmarshal the response:
	response := new(CollectionsGetTreeResponse)
	if err := json.Unmarshal([]byte(rspJSON), response); err != nil {
		return nil, fmt.Errorf("Error unmarshaling response: %w", err)
	}

	return response, nil
}
