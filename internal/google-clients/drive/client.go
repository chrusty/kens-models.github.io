package googledriveclient

import (
	"context"
	"encoding/base64"
	"fmt"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/option"
)

type Client struct {
	service *drive.Service
}

func New() *Client {
	return new(Client)
}

func (c *Client) WithAPIKey(apiKey string) (*Client, error) {

	// Read credentials from an env-var:
	credBytes, err := base64.StdEncoding.DecodeString(apiKey)
	if err != nil {
		return c, fmt.Errorf("Unable to decode Google auth key: %w", err)
	}

	// Authenticate and get configuration:
	config, err := google.JWTConfigFromJSON(credBytes, drive.DriveScope, drive.DriveReadonlyScope, drive.DriveMetadataScope, drive.DriveMetadataReadonlyScope)
	if err != nil {
		return c, fmt.Errorf("Unable to get JWT config: %w", err)
	}

	// Prepare a spreadsheet service:
	service, err := drive.NewService(context.TODO(), option.WithHTTPClient(config.Client(context.TODO())))
	if err != nil {
		return c, fmt.Errorf("Unable to prepare service: %w", err)
	}

	c.service = service

	return c, nil
}

func (c *Client) Service() *drive.Service {
	return c.service
}
