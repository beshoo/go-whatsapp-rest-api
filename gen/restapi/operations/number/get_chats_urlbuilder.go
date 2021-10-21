// Code generated by go-swagger; DO NOT EDIT.

package number

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

// GetChatsURL generates an URL for the get chats operation
type GetChatsURL struct {
	PhoneNumber string

	BeforeMessageID  *string
	FromMe           *bool
	NumberOfMessages int64
	SessionID        strfmt.UUID4

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetChatsURL) WithBasePath(bp string) *GetChatsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetChatsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetChatsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/number/{phoneNumber}/chats"

	phoneNumber := o.PhoneNumber
	if phoneNumber != "" {
		_path = strings.Replace(_path, "{phoneNumber}", phoneNumber, -1)
	} else {
		return nil, errors.New("phoneNumber is required on GetChatsURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var beforeMessageIDQ string
	if o.BeforeMessageID != nil {
		beforeMessageIDQ = *o.BeforeMessageID
	}
	if beforeMessageIDQ != "" {
		qs.Set("beforeMessageId", beforeMessageIDQ)
	}

	var fromMeQ string
	if o.FromMe != nil {
		fromMeQ = swag.FormatBool(*o.FromMe)
	}
	if fromMeQ != "" {
		qs.Set("fromMe", fromMeQ)
	}

	numberOfMessagesQ := swag.FormatInt64(o.NumberOfMessages)
	if numberOfMessagesQ != "" {
		qs.Set("numberOfMessages", numberOfMessagesQ)
	}

	sessionIDQ := o.SessionID.String()
	if sessionIDQ != "" {
		qs.Set("sessionId", sessionIDQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetChatsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetChatsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetChatsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetChatsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetChatsURL")
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
func (o *GetChatsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
