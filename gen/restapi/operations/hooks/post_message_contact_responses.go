// Code generated by go-swagger; DO NOT EDIT.

package hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostMessageContactOKCode is the HTTP code returned for type PostMessageContactOK
const PostMessageContactOKCode int = 200

/*PostMessageContactOK Return 200 else the api will retry

swagger:response postMessageContactOK
*/
type PostMessageContactOK struct {
}

// NewPostMessageContactOK creates PostMessageContactOK with default headers values
func NewPostMessageContactOK() *PostMessageContactOK {

	return &PostMessageContactOK{}
}

// WriteResponse to the client
func (o *PostMessageContactOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
