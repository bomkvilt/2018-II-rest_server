// Code generated by go-swagger; DO NOT EDIT.

package forum

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// ForumGetOneURL generates an URL for the forum get one operation
type ForumGetOneURL struct {
	Slug string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ForumGetOneURL) WithBasePath(bp string) *ForumGetOneURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ForumGetOneURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ForumGetOneURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/forum/{slug}/details"

	slug := o.Slug
	if slug != "" {
		_path = strings.Replace(_path, "{slug}", slug, -1)
	} else {
		return nil, errors.New("Slug is required on ForumGetOneURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ForumGetOneURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ForumGetOneURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ForumGetOneURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ForumGetOneURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ForumGetOneURL")
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
func (o *ForumGetOneURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
