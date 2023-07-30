package sheetsdata

import (
	sheetsClient "github.com/chrusty/kens-models.github.io/internal/google-clients/sheets"
)

func ValuesGet(client *sheetsClient.Client, spreadsheetId string, selection string, header bool) (*Values, error) {

	// Retrieve rows:
	valuesGetResponse, err := client.Service().Spreadsheets.Values.Get(spreadsheetId, selection).Do()
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
