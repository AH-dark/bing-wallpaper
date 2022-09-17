package bing

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	*http.Client
}

func NewClient() *Client {
	return &Client{
		Client: http.DefaultClient,
	}
}

func (c *Client) Fetch(day, num int) (*WallpaperData, *http.Response, error) {
	resp, err := c.Get(fmt.Sprintf("https://cn.bing.com/HPImageArchive.aspx?format=js&idx=%d&n=%d", day, num))
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	var data WallpaperData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, nil, err
	}

	return &data, resp, nil
}
