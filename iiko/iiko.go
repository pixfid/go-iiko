/*
 * MIT License
 *
 * Copyright (c) 2020 Semchenko Aleksandr
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package iiko

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
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

type ErrorOut struct {
	Error Error
}

type Error struct {
	Code     int
	Body     string
	Endpoint string
}

func (e Error) Error() string {
	return fmt.Sprintf(`Status Code: %d`, e.Code)
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 10 * time.Second}
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

	var req *http.Request

	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	switch method {
	case "GET":
		req, err = http.NewRequest(method, u.String(), nil)
		if err != nil {
			return nil, err
		}
	case "POST":
		postData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, u.String(), bytes.NewBuffer(postData))
		if err != nil {
			return nil, err
		}
	default:

	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	if c.UserAgent == "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, target interface{}) ([]byte, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}
	req = withContext(ctx, req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 || resp.StatusCode == 201 {

		if isPlainTextResponse(b) {
			return b, err
		}

		err = json.Unmarshal(b, target)
		if err != nil {
			return nil, err
		}
		return b, nil
	}

	err = Error{
		Code:     resp.StatusCode,
		Body:     string(b),
		Endpoint: req.URL.String(),
	}
	return nil, err
}

func isPlainTextResponse(b []byte) bool {
	return bytes.HasPrefix(b, []byte(`"`))
}

func withContext(ctx context.Context, req *http.Request) *http.Request {
	return req.WithContext(ctx)
}
