package flickrclient

const (
	argFormat                = "format"
	argFormatJSON            = "json"
	argFormatXML             = "xml"
	argNoJSONCallBack        = "nojsoncallback"
	argUserId                = "user_id"
	methodCollectionsGetTree = "flickr.collections.getTree"
)

type Client struct {
	apiKey string
	format string
	userId string
	secret string
}

func New(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		format: argFormatJSON,
	}
}

func (c *Client) WithSecret(secret string) *Client {
	c.secret = secret
	return c
}

func (c *Client) WithUserId(userId string) *Client {
	c.userId = userId
	return c
}
