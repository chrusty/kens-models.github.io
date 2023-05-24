package flickrclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCollectionsJSON = `{
	"collections": {
		"collection": [
		{
			"id": "1",
			"title": "Collection 1",
			"description": "This is the first collection",
			"iconlarge": "https:\\/\\/combo.staticflickr.com\\/pw\\/images\\/collection_default_l.gif",
			"iconsmall": "https:\\/\\/combo.staticflickr.com\\/pw\\/images\\/collection_default_s.gif",
			"set": [
			{
				"id": "11",
				"title": "Album 1.1",
				"description": "This is the first album in the first collection."
			},
			{
				"id": "12",
				"title": "Album 1.2",
				"description": "This is the second album in the first collection."
			}
			]
		},
		{
			"id": "2",
			"title": "Collection 2",
			"description": "This is the second collection",
			"iconlarge": "https:\\/\\/combo.staticflickr.com\\/pw\\/images\\/collection_default_l.gif",
			"iconsmall": "https:\\/\\/combo.staticflickr.com\\/pw\\/images\\/collection_default_s.gif",
			"set": [
			{
				"id": "21",
				"title": "Album 2.1",
				"description": "This is the first album in the second collection."
			},
			{
				"id": "22",
				"title": "Album 2.2",
				"description": "This is the second album in the second collection."
			}
			]
		}
		]
	},
	"stat": "ok"
}`

func TestCollections(t *testing.T) {

	// Unmarshal JSON into collections:
	rsp := new(CollectionsGetTreeResponse)
	assert.NoError(t, json.Unmarshal([]byte(testCollectionsJSON), rsp))
	assert.Equal(t, "ok", rsp.Status)
	assert.Len(t, rsp.Collections.Collection, 2)
	assert.Equal(t, rsp.Collections.Collection[0].ID, "1")
	assert.Len(t, rsp.Collections.Collection[0].Set, 2)
	assert.Equal(t, rsp.Collections.Collection[0].Set[0].ID, "11")
}
