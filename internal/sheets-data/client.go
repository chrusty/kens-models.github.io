package sheetsdata

import (
	"context"
	"encoding/base64"
	"fmt"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const (
	sheetsAPIAddress = "https://www.googleapis.com/auth/spreadsheets"
)

type Client struct {
	sheetsService *sheets.Service
	spreadsheetId string
}

func New(spreadsheetId string) *Client {
	return &Client{spreadsheetId: spreadsheetId}
}

func (c *Client) WithAPIKey(apiKey string) (*Client, error) {

	// Read credentials from an env-var:
	credBytes, err := base64.StdEncoding.DecodeString(apiKey)
	if err != nil {
		return c, fmt.Errorf("Unable to decode Google auth key: %w", err)
	}

	// Authenticate and get configuration:
	config, err := google.JWTConfigFromJSON(credBytes, sheetsAPIAddress)
	if err != nil {
		return c, fmt.Errorf("Unable to get JWT config: %w", err)
	}

	// Prepare a spreadsheet service:
	service, err := sheets.NewService(context.TODO(), option.WithHTTPClient(config.Client(context.TODO())))
	if err != nil {
		return c, fmt.Errorf("Unable to prepare sheets service: %w", err)
	}

	c.sheetsService = service

	return c, nil
}

func (c *Client) Service() *sheets.Service {
	return c.sheetsService
}
