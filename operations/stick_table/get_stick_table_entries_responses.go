// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package stick_table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// GetStickTableEntriesOKCode is the HTTP code returned for type GetStickTableEntriesOK
const GetStickTableEntriesOKCode int = 200

/*GetStickTableEntriesOK Successful operation

swagger:response getStickTableEntriesOK
*/
type GetStickTableEntriesOK struct {

	/*
	  In: Body
	*/
	Payload models.StickTableEntries `json:"body,omitempty"`
}

// NewGetStickTableEntriesOK creates GetStickTableEntriesOK with default headers values
func NewGetStickTableEntriesOK() *GetStickTableEntriesOK {

	return &GetStickTableEntriesOK{}
}

// WithPayload adds the payload to the get stick table entries o k response
func (o *GetStickTableEntriesOK) WithPayload(payload models.StickTableEntries) *GetStickTableEntriesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stick table entries o k response
func (o *GetStickTableEntriesOK) SetPayload(payload models.StickTableEntries) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStickTableEntriesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.StickTableEntries{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetStickTableEntriesDefault General Error

swagger:response getStickTableEntriesDefault
*/
type GetStickTableEntriesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetStickTableEntriesDefault creates GetStickTableEntriesDefault with default headers values
func NewGetStickTableEntriesDefault(code int) *GetStickTableEntriesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetStickTableEntriesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get stick table entries default response
func (o *GetStickTableEntriesDefault) WithStatusCode(code int) *GetStickTableEntriesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get stick table entries default response
func (o *GetStickTableEntriesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get stick table entries default response
func (o *GetStickTableEntriesDefault) WithConfigurationVersion(configurationVersion int64) *GetStickTableEntriesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get stick table entries default response
func (o *GetStickTableEntriesDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get stick table entries default response
func (o *GetStickTableEntriesDefault) WithPayload(payload *models.Error) *GetStickTableEntriesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stick table entries default response
func (o *GetStickTableEntriesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStickTableEntriesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
