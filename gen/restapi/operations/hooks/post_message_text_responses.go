// Code generated by go-swagger; DO NOT EDIT.

package hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostMessageTextOKCode is the HTTP code returned for type PostMessageTextOK
const PostMessageTextOKCode int = 200

/*PostMessageTextOK Return 200 else the api will retry

swagger:response postMessageTextOK
*/
type PostMessageTextOK struct {
}

// NewPostMessageTextOK creates PostMessageTextOK with default headers values
func NewPostMessageTextOK() *PostMessageTextOK {

	return &PostMessageTextOK{}
}

// WriteResponse to the client
func (o *PostMessageTextOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
