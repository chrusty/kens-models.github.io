package sheetsdata

import "fmt"

type Values struct {
	Header  []string
	Rows    [][]interface{}
	Entries map[string]map[string]string
}

func (v *Values) CreateMap(key string) error {

	// Find which header field is the key:
	var headerKey int = 999999
	for i, header := range v.Header {
		if header == key {
			headerKey = i
			break
		}
	}
	if headerKey == 999999 {
		return fmt.Errorf("Header key %s not found", key)
	}

	// Establish the data structure:
	v.Entries = make(map[string]map[string]string)

	// Put the rows into an MSI:
	for _, row := range v.Rows {
		if len(row) < headerKey {
			continue
		}
		rowKey := row[headerKey].(string)
		rowData := make(map[string]string)
		for i, header := range v.Header {
			if len(row) <= i {
				break
			}
			rowData[header] = row[i].(string)
		}
		v.Entries[rowKey] = rowData
	}

	return nil
}

func (v *Values) GetRow(key string) (map[string]string, error) {

	// Look up the row:
	row, exists := v.Entries[key]
	if !exists {
		return nil, fmt.Errorf("Row with key %s not found", key)
	}

	return row, nil
}
