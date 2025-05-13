package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
)

var (
	ErrNilClient = errors.New("given HTTP client is nil")
)

// HTTPClient is the minimum implementation of an HTTP client to be used
// by the IEEEClient.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

var _ HTTPClient = (*http.Client)(nil)

// IEEEClient is an HTTP client specialized at handling IEEE API specificities.
// It leverages the client passed at NewIEEEClient to mimic a valid IEEE API
// behavior.
type IEEEClient struct {
	sub HTTPClient
}

// NewIEEEClient builds a new *IEEEClient given the subsidiary HTTP client.
// If the given client is nil, returns a nil *IEEEClient and an error.
func NewIEEEClient(sub HTTPClient) (*IEEEClient, error) {
	if sub == nil {
		return nil, ErrNilClient
	}
	return &IEEEClient{
		sub: sub,
	}, nil
}

func (client *IEEEClient) do(req *http.Request, opts ...Option) (*http.Response, error) {
	// Based on the Mozilla display server, enough to fake a valid API call
	req.Header.Add("User-Agent", "X11")
	// Make sure IEEE API sees it as a valid call
	req.Header.Set("Referer", "https://ieeexplore.ieee.org")

	// Apply functional options
	req = applyOpts(req, opts...)

	res, err := client.sub.Do(req)
	if err != nil {
		return nil, err
	}

	// Check status code, should always be a 200 OK
	if res.StatusCode != http.StatusOK {
		return nil, &ErrUnexpectedStatus{
			StatusCode: res.StatusCode,
		}
	}

	return res, nil
}

// Option is the generic interface for this API functional options.
type Option interface {
	apply(*options)
}

type options struct {
	Ctx context.Context
}

type ctxOption struct {
	Ctx context.Context
}

func (opt ctxOption) apply(opts *options) {
	opts.Ctx = opt.Ctx
}

// WithContext enable specifying a context to pass to the client issuing the
// API call.
func WithContext(ctx context.Context) Option {
	return &ctxOption{
		Ctx: ctx,
	}
}

func applyOpts(req *http.Request, opts ...Option) *http.Request {
	reqopts := &options{
		Ctx: context.Background(),
	}

	for _, opt := range opts {
		opt.apply(reqopts)
	}

	req = req.WithContext(reqopts.Ctx)

	return req
}

// ErrUnexpectedStatus is returned when the IEEE API gave back an unexpected
// status code, meaning something went wrong.
type ErrUnexpectedStatus struct {
	StatusCode int
}

func (e ErrUnexpectedStatus) Error() string {
	return fmt.Sprintf("got unexpected status code %d", e.StatusCode)
}

var _ error = (*ErrUnexpectedStatus)(nil)

func (client *IEEEClient) get(endp string, params, dst any, opts ...Option) error {
	// Build the API call
	tgt := fmt.Sprintf("https://ieeexplore.ieee.org/rest/%s", endp)
	req, _ := http.NewRequest(http.MethodGet, tgt, nil)

	// Encode URL parameters if any
	if params != nil {
		values := url.Values{}
		if err := schema.NewEncoder().Encode(params, values); err != nil {
			return err
		}
		req.URL.RawQuery = values.Encode()
	}

	// Issue call
	res, err := client.do(req, opts...)
	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	// Decode JSON REST response
	return json.NewDecoder(res.Body).Decode(dst)
}

func (client *IEEEClient) post(endp string, params, dst any, opts ...Option) error {
	// Build JSON content
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(params); err != nil {
		return err
	}

	// Build the API call
	tgt := fmt.Sprintf("https://ieeexplore.ieee.org/rest/%s", endp)
	req, _ := http.NewRequest(http.MethodPost, tgt, buf)
	req.Header.Set("Content-Type", "application/json")

	// Issue call
	res, err := client.do(req, opts...)
	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	// Decode JSON REST response
	return json.NewDecoder(res.Body).Decode(dst)
}

type UserInfo struct {
	Institute                     bool    `json:"institute"`
	Member                        bool    `json:"member"`
	Individual                    bool    `json:"individual"`
	Guest                         bool    `json:"guest"`
	SubscribedContent             bool    `json:"subscribedContent"`
	FileCabinetContent            bool    `json:"fileCabinetContent"`
	FileCabinetUser               bool    `json:"fileCabinetUser"`
	InstitutionalFileCabinetUser  bool    `json:"institutionalFileCabinetUser"`
	IP                            *string `json:"ip,omitempty"`
	ShowPatentCitations           bool    `json:"showPatentCitations"`
	ShowGet802Link                bool    `json:"showGet802Link"`
	ShowOpenURLLink               *bool   `json:"showOpenUrlLink,omitempty"`
	Tracked                       *bool   `json:"tracked,omitempty"`
	DelegatedAdmin                *bool   `json:"delegatedAdmin,omitempty"`
	Desktop                       *bool   `json:"desktop,omitempty"`
	IsInstitutionDashboardEnabled *bool   `json:"isInstitutionDashboardEnabled,omitempty"`
	IsInstitutionProfileEnabled   *bool   `json:"isInstitutionProfileEnabled,omitempty"`
	IsRoamingEnabled              *bool   `json:"isRoamingEnabled,omitempty"`
	IsDelegatedAdmin              *bool   `json:"isDelegatedAdmin,omitempty"`
	IsMdl                         *bool   `json:"isMdl,omitempty"`
	IsCwg                         *bool   `json:"isCwg,omitempty"`
	IsAcademic                    *bool   `json:"isAcademic,omitempty"`
	IsIel                         *bool   `json:"isIel,omitempty"`
	IsReadAndPublish              *bool   `json:"isReadAndPublish,omitempty"`
}
