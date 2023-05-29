package sheetsdata

func (c *Client) ValuesGet(selection string, header bool) (*Values, error) {

	// Retrieve rows:
	valuesGetResponse, err := c.sheetsService.Spreadsheets.Values.Get(c.spreadsheetId, selection).Do()
	if err != nil || valuesGetResponse.HTTPStatusCode != 200 {
		return nil, err
	}

	// Put the response into a set of values:
	response := new(Values)
	if header {
		for _, header := range valuesGetResponse.Values[0] {
			response.Header = append(response.Header, header.(string))
		}
		response.Rows = valuesGetResponse.Values[1:]
	} else {
		response.Rows = valuesGetResponse.Values
	}

	return response, nil
}
