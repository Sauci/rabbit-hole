package rabbithole

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type PublishInfo struct {
	Properties      *Properties `json:"properties"`
	RoutingKey      string      `json:"routing_key"`
	Payload         string      `json:"payload"`
	PayloadEncoding string      `json:"payload_encoding"`
}

func (c *Client) PostPublish(vhostname string, exchange string, message PublishInfo) (res *http.Response, err error) {
	body, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	req, err := newRequestWithBody(c, "POST", "exchanges/"+url.PathEscape(vhostname)+"/"+exchange+"/publish", body)
	if err != nil {
		return nil, err
	}

	if res, err = executeRequest(c, req); err != nil {
		return nil, err
	}

	return res, nil
}
