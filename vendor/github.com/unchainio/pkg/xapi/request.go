package xapi

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/unchainio/pkg/errors"
)

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is encoded by the client's encoder and included as the
// request body.
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		c.BaseURL.Path += "/"
	}

	u, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	var bodyReader io.Reader

	if r, ok := body.(io.Reader); ok {
		bodyReader = r
	} else if body != nil {
		bodyReader, err = c.opts.encoder.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/"+c.opts.encoder.Type())
	}
	req.Header.Set("Accept", "application/"+c.opts.encoder.Type())
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

type CleanupFn func()

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, CleanupFn, error) {
	req = req.WithContext(ctx)

	resp, err := c.Client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = SanitizeURL(url).String()
				return nil, nil, e
			}
		}

		return nil, nil, errors.Wrap(err, "")
	}

	cleanup := func() {
		resp.Body.Close()
	}

	if c := resp.StatusCode; 200 > c || c > 299 {
		b, _ := ioutil.ReadAll(resp.Body)
		cleanup()

		return resp, nil, errors.Errorf("(%d) %s: %s", c, http.StatusText(c), string(b))
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := c.opts.encoder.Decode(resp.Body, v)

			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = errors.Wrap(decErr, "")
			}
		}
	}

	return resp, cleanup, err
}

// SanitizeURL redacts the client_secret parameter from the URL which may be
// exposed to the user.
func SanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

func (c *Client) SimpleRequest(ctx context.Context, method string, path string, input interface{}, output interface{}) (*http.Request, *http.Response, func(), error) {
	req, err := c.NewRequest(method, path, input)

	if err != nil {
		return nil, nil, nil, err
	}

	res, cleanup, err := c.Do(ctx, req, output)

	if err != nil {
		return nil, nil, nil, err
	}

	return req, res, cleanup, nil
}
