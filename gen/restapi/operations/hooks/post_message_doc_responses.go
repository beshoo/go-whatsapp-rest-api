// Code generated by go-swagger; DO NOT EDIT.

package hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostMessageDocOKCode is the HTTP code returned for type PostMessageDocOK
const PostMessageDocOKCode int = 200

/*PostMessageDocOK Return 200 else the api will retry

swagger:response postMessageDocOK
*/
type PostMessageDocOK struct {
}

// NewPostMessageDocOK creates PostMessageDocOK with default headers values
func NewPostMessageDocOK() *PostMessageDocOK {

	return &PostMessageDocOK{}
}

// WriteResponse to the client
func (o *PostMessageDocOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
