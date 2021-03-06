// Code generated by go-swagger; DO NOT EDIT.

package hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostMessageLocationOKCode is the HTTP code returned for type PostMessageLocationOK
const PostMessageLocationOKCode int = 200

/*PostMessageLocationOK Return 200 else the api will retry

swagger:response postMessageLocationOK
*/
type PostMessageLocationOK struct {
}

// NewPostMessageLocationOK creates PostMessageLocationOK with default headers values
func NewPostMessageLocationOK() *PostMessageLocationOK {

	return &PostMessageLocationOK{}
}

// WriteResponse to the client
func (o *PostMessageLocationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
