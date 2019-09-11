package xapi

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Client struct {
	BaseURL   *url.URL
	Client    *http.Client
	UserAgent string

	opts *Options
}

type Options struct {
	encoder   Encoder
	userAgent string
	client    *http.Client
}

type OptionFunc func(*Options)

// NewClient returns a new generic API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(baseURLString string, optFns ...OptionFunc) (*Client, error) {
	opts := &Options{
		// TODO(e-nikolov) Does it matter that encoder could be allocated twice?
		encoder:   NewJSONEncoder(),
		userAgent: "xapi",
		client:    http.DefaultClient,
	}

	for _, optFn := range optFns {
		optFn(opts)
	}

	baseURL, err := url.Parse(baseURLString)

	if err != nil {
		return nil, errors.Wrap(err, "Could not parse url from the config")
	}

	return &Client{
		BaseURL:   baseURL,
		Client:    opts.client,
		UserAgent: opts.userAgent,

		opts: opts,
	}, nil
}

func WithClient(client *http.Client) OptionFunc {
	return func(o *Options) {
		o.client = client
	}
}

func XML() OptionFunc {
	return func(o *Options) {
		o.encoder = NewXMLEncoder()
	}
}

func JSON() OptionFunc {
	return func(o *Options) {
		o.encoder = NewJSONEncoder()
	}
}

func WithEncoder(encoder Encoder) OptionFunc {
	return func(o *Options) {
		o.encoder = encoder
	}
}

func WithUserAgent(userAgent string) OptionFunc {
	return func(o *Options) {
		o.userAgent = userAgent
	}
}
