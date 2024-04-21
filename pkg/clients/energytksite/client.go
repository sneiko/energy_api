package energytksite

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const baseUrl = "https://nrg-tk.ru/rest"

type Client struct {
	client *resty.Client
}

func New() *Client {
	client := resty.New().
		SetBaseURL(baseUrl).
		SetRetryCount(3).
		SetRetryWaitTime(5 * time.Second)
	return &Client{
		client: client,
	}
}

func (c *Client) CheckInvoice(number string) (*InvoiceResult, error) {
	var response InvoiceResult
	res, err := c.client.R().
		SetResult(&response).
		SetHeader("Content-Type", "application/json;charset=utf-8").
		Get("/tracking/?docNum=" + number)
	if err != nil {
		return nil, err
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf("error: %s", res.String())
	}

	return &response, nil
}
