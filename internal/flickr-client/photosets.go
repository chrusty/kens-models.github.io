package flickrclient

import (
	"encoding/json"
	"fmt"

	"github.com/mncaudill/go-flickr"
)

type responseContent struct {
	Content string `json:"_content"`
}

func (r responseContent) String() string {
	return r.Content
}

type primaryPhotoExtras struct {
	UrlSquare      string `json:"url_sq"`
	HeightSquare   int    `json:"height_sq"`
	WidthSquare    int    `json:"width_sq"`
	UrlThumb       string `json:"url_t"`
	HeightThumb    int    `json:"height_t"`
	WidthThumb     int    `json:"width_t"`
	UrlSmall       string `json:"url_s"`
	HeightSmall    int    `json:"height_s"`
	WidthSmall     int    `json:"width_s"`
	UrlMedium      string `json:"url_m"`
	HeightMedium   int    `json:"height_m"`
	WidthMedium    int    `json:"width_m"`
	UrlOriginal    string `json:"url_o"`
	HeightOriginal int    `json:"height_o"`
	WidthOriginal  int    `json:"width_o"`
}

type Photoset struct {
	ID                  string             `json:"id"`
	Owner               string             `json:"iownerd"`
	Username            string             `json:"username"`
	Primary             string             `json:"primary"`
	Secret              string             `json:"secret"`
	Server              string             `json:"server"`
	Farm                int                `json:"farm"`
	CountViews          string             `json:"count_views"`
	CountComments       string             `json:"count_comments"`
	CountPhotos         int                `json:"count_photos"`
	CountVideos         int                `json:"count_videos"`
	Title               responseContent    `json:"title"`
	Description         responseContent    `json:"description"`
	CanComment          int                `json:"can_comment"`
	DateCreate          string             `json:"date_create"`
	DateUpdate          string             `json:"date_update"`
	Photos              int                `json:"photos"`
	Videos              int                `json:"videos"`
	VisibilityCanSeeSet int                `json:"visibility_can_see_set"`
	NeedsInterstitial   int                `json:"needs_interstitial"`
	PrimaryPhotoExtras  primaryPhotoExtras `json:"primary_photo_extras"`
}

type Photosets struct {
	Page     int        `json:"page"`
	Pages    int        `json:"pages"`
	Perpage  int        `json:"perpage"`
	Total    int        `json:"total"`
	Photoset []Photoset `json:"photoset"`
}

type PhotosetsGetListResponse struct {
	Photosets Photosets `json:"photosets"`
	Status    string    `json:"stat"`
}

func (c *Client) PhotosetsGetList() (*PhotosetsGetListResponse, error) {

	// Build the request:
	req := &flickr.Request{
		ApiKey: c.apiKey,
		Method: methodPhotosetsGetList,
		Args: map[string]string{
			argUserId:             c.userId,
			argFormat:             c.format,
			argNoJSONCallBack:     "1",
			argPrimaryPhotoExtras: "url_sq,url_t,url_s,url_m,url_o",
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
	response := new(PhotosetsGetListResponse)
	if err := json.Unmarshal([]byte(rspJSON), response); err != nil {
		return nil, fmt.Errorf("Error unmarshaling response: %w", err)
	}

	return response, nil
}
