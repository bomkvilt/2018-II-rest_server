// Code generated by go-swagger; DO NOT EDIT.

package forum

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ForumGetThreadsURL generates an URL for the forum get threads operation
type ForumGetThreadsURL struct {
	Slug string

	Desc  *bool
	Limit *int32
	Since *strfmt.DateTime

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ForumGetThreadsURL) WithBasePath(bp string) *ForumGetThreadsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ForumGetThreadsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ForumGetThreadsURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/forum/{slug}/threads"

	slug := o.Slug
	if slug != "" {
		_path = strings.Replace(_path, "{slug}", slug, -1)
	} else {
		return nil, errors.New("Slug is required on ForumGetThreadsURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var desc string
	if o.Desc != nil {
		desc = swag.FormatBool(*o.Desc)
	}
	if desc != "" {
		qs.Set("desc", desc)
	}

	var limit string
	if o.Limit != nil {
		limit = swag.FormatInt32(*o.Limit)
	}
	if limit != "" {
		qs.Set("limit", limit)
	}

	var since string
	if o.Since != nil {
		since = o.Since.String()
	}
	if since != "" {
		qs.Set("since", since)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ForumGetThreadsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ForumGetThreadsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ForumGetThreadsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ForumGetThreadsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ForumGetThreadsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ForumGetThreadsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
