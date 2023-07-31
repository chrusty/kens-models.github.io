package googlesheetsclient

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Client struct {
	http    *http.Client
	service *sheets.Service
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
	config, err := google.JWTConfigFromJSON(credBytes, sheets.SpreadsheetsReadonlyScope)
	if err != nil {
		return c, fmt.Errorf("Unable to get JWT config: %w", err)
	}

	// Prepare a spreadsheet service:
	service, err := sheets.NewService(context.TODO(), option.WithHTTPClient(config.Client(context.TODO())))
	if err != nil {
		return c, fmt.Errorf("Unable to prepare service: %w", err)
	}

	c.http = config.Client(context.TODO())
	c.service = service

	return c, nil
}

func (c *Client) HTTP() *http.Client {
	return c.http
}

func (c *Client) Service() *sheets.Service {
	return c.service
}
