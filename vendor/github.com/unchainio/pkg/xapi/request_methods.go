package xapi

import (
	"context"
	"net/http"
)

func (c *Client) Get(ctx context.Context, path string, input interface{}, output interface{}) (*http.Request, *http.Response, func(), error) {
	return c.SimpleRequest(ctx, http.MethodGet, path, input, output)
}

func (c *Client) Post(ctx context.Context, path string, input interface{}, output interface{}) (*http.Request, *http.Response, func(), error) {
	return c.SimpleRequest(ctx, http.MethodPost, path, input, output)
}

func (c *Client) Patch(ctx context.Context, path string, input interface{}, output interface{}) (*http.Request, *http.Response, func(), error) {
	return c.SimpleRequest(ctx, http.MethodPatch, path, input, output)
}

func (c *Client) Put(ctx context.Context, path string, input interface{}, output interface{}) (*http.Request, *http.Response, func(), error) {
	return c.SimpleRequest(ctx, http.MethodPut, path, input, output)
}

func (c *Client) Delete(ctx context.Context, path string, input interface{}, output interface{}) (*http.Request, *http.Response, func(), error) {
	return c.SimpleRequest(ctx, http.MethodDelete, path, input, output)
}

func (c *Client) Trace(ctx context.Context, path string, input interface{}, output interface{}) (*http.Request, *http.Response, func(), error) {
	return c.SimpleRequest(ctx, http.MethodTrace, path, input, output)
}

func (c *Client) Connect(ctx context.Context, path string, input interface{}, output interface{}) (*http.Request, *http.Response, func(), error) {
	return c.SimpleRequest(ctx, http.MethodConnect, path, input, output)
}

func (c *Client) Head(ctx context.Context, path string, input interface{}, output interface{}) (*http.Request, *http.Response, func(), error) {
	return c.SimpleRequest(ctx, http.MethodHead, path, input, output)
}

func (c *Client) Options(ctx context.Context, path string, input interface{}, output interface{}) (*http.Request, *http.Response, func(), error) {
	return c.SimpleRequest(ctx, http.MethodOptions, path, input, output)
}
