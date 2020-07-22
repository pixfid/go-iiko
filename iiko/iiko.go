package iiko

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://iiko.biz:9900/api/0/"
	userAgent      = "go-iiko"
)

type Client struct {
	BaseURL      *url.URL
	client       *http.Client
	common       service
	Server       string
	Auth         *AuthService
	Orders       *Orders
	RmsSettings  *RmsSettings
	Nomenclature *Nomenclature
	Cities       *CitiesService
	StopLists    *StopLists
	UserAgent    string
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.common.client = c
	c.Auth = (*AuthService)(&c.common)
	c.Orders = (*Orders)(&c.common)
	c.Nomenclature = (*Nomenclature)(&c.common)
	c.Cities = (*CitiesService)(&c.common)
	c.StopLists = (*StopLists)(&c.common)
	c.RmsSettings = (*RmsSettings)(&c.common)
	return c
}

func (c *Client) NewRequest(method string, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) ([]byte, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}
	req = withContext(ctx, req)

	response, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	defer response.Body.Close()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, response.Body)
		} else {
			decErr := json.NewDecoder(response.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func withContext(ctx context.Context, req *http.Request) *http.Request {
	return req.WithContext(ctx)
}
