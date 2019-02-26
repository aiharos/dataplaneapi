// Code generated by go-swagger; DO NOT EDIT.

package bind

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// DeleteBindAcceptedCode is the HTTP code returned for type DeleteBindAccepted
const DeleteBindAcceptedCode int = 202

/*DeleteBindAccepted Configuration change accepted and reload requested

swagger:response deleteBindAccepted
*/
type DeleteBindAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteBindAccepted creates DeleteBindAccepted with default headers values
func NewDeleteBindAccepted() *DeleteBindAccepted {

	return &DeleteBindAccepted{}
}

// WithReloadID adds the reloadId to the delete bind accepted response
func (o *DeleteBindAccepted) WithReloadID(reloadID string) *DeleteBindAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete bind accepted response
func (o *DeleteBindAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteBindAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteBindNoContentCode is the HTTP code returned for type DeleteBindNoContent
const DeleteBindNoContentCode int = 204

/*DeleteBindNoContent Bind deleted

swagger:response deleteBindNoContent
*/
type DeleteBindNoContent struct {
}

// NewDeleteBindNoContent creates DeleteBindNoContent with default headers values
func NewDeleteBindNoContent() *DeleteBindNoContent {

	return &DeleteBindNoContent{}
}

// WriteResponse to the client
func (o *DeleteBindNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteBindNotFoundCode is the HTTP code returned for type DeleteBindNotFound
const DeleteBindNotFoundCode int = 404

/*DeleteBindNotFound The specified resource was not found

swagger:response deleteBindNotFound
*/
type DeleteBindNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteBindNotFound creates DeleteBindNotFound with default headers values
func NewDeleteBindNotFound() *DeleteBindNotFound {

	return &DeleteBindNotFound{}
}

// WithPayload adds the payload to the delete bind not found response
func (o *DeleteBindNotFound) WithPayload(payload *models.Error) *DeleteBindNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete bind not found response
func (o *DeleteBindNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBindNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteBindDefault General Error

swagger:response deleteBindDefault
*/
type DeleteBindDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteBindDefault creates DeleteBindDefault with default headers values
func NewDeleteBindDefault(code int) *DeleteBindDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteBindDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete bind default response
func (o *DeleteBindDefault) WithStatusCode(code int) *DeleteBindDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete bind default response
func (o *DeleteBindDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete bind default response
func (o *DeleteBindDefault) WithPayload(payload *models.Error) *DeleteBindDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete bind default response
func (o *DeleteBindDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBindDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}