package xapi

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func (c *Client) Upload(ctx context.Context, path string, readers map[string]io.Reader, v interface{}) (*http.Request, *http.Response, func(), error) {
	var err error

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	closeFn := func(r io.Reader) {
		if r, ok := r.(io.Closer); ok {
			r.Close()
		}
	}

	for key, r := range readers {
		var fw io.Writer

		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				closeFn(r)
				return nil, nil, nil, err
			}
		} else {
			if fw, err = w.CreateFormFile(key, key); err != nil {
				closeFn(r)
				return nil, nil, nil, err
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			closeFn(r)
			return nil, nil, nil, err
		}

		closeFn(r)
	}
	w.Close()

	req, err := c.NewRequest(http.MethodPost, path, &b)
	if err != nil {
		return nil, nil, nil, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	res, cleanup, err := c.Do(ctx, req, v)

	if err != nil {
		return nil, nil, nil, err
	}

	return req, res, cleanup, nil
}
