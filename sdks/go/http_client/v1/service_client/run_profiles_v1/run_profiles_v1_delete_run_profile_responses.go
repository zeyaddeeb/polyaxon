// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package run_profiles_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// RunProfilesV1DeleteRunProfileReader is a Reader for the RunProfilesV1DeleteRunProfile structure.
type RunProfilesV1DeleteRunProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunProfilesV1DeleteRunProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunProfilesV1DeleteRunProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewRunProfilesV1DeleteRunProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRunProfilesV1DeleteRunProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunProfilesV1DeleteRunProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRunProfilesV1DeleteRunProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunProfilesV1DeleteRunProfileOK creates a RunProfilesV1DeleteRunProfileOK with default headers values
func NewRunProfilesV1DeleteRunProfileOK() *RunProfilesV1DeleteRunProfileOK {
	return &RunProfilesV1DeleteRunProfileOK{}
}

/*RunProfilesV1DeleteRunProfileOK handles this case with default header values.

A successful response.
*/
type RunProfilesV1DeleteRunProfileOK struct {
}

func (o *RunProfilesV1DeleteRunProfileOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/run_profiles/{uuid}][%d] runProfilesV1DeleteRunProfileOK ", 200)
}

func (o *RunProfilesV1DeleteRunProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRunProfilesV1DeleteRunProfileNoContent creates a RunProfilesV1DeleteRunProfileNoContent with default headers values
func NewRunProfilesV1DeleteRunProfileNoContent() *RunProfilesV1DeleteRunProfileNoContent {
	return &RunProfilesV1DeleteRunProfileNoContent{}
}

/*RunProfilesV1DeleteRunProfileNoContent handles this case with default header values.

No content.
*/
type RunProfilesV1DeleteRunProfileNoContent struct {
	Payload interface{}
}

func (o *RunProfilesV1DeleteRunProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/run_profiles/{uuid}][%d] runProfilesV1DeleteRunProfileNoContent  %+v", 204, o.Payload)
}

func (o *RunProfilesV1DeleteRunProfileNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *RunProfilesV1DeleteRunProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunProfilesV1DeleteRunProfileForbidden creates a RunProfilesV1DeleteRunProfileForbidden with default headers values
func NewRunProfilesV1DeleteRunProfileForbidden() *RunProfilesV1DeleteRunProfileForbidden {
	return &RunProfilesV1DeleteRunProfileForbidden{}
}

/*RunProfilesV1DeleteRunProfileForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type RunProfilesV1DeleteRunProfileForbidden struct {
	Payload interface{}
}

func (o *RunProfilesV1DeleteRunProfileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/run_profiles/{uuid}][%d] runProfilesV1DeleteRunProfileForbidden  %+v", 403, o.Payload)
}

func (o *RunProfilesV1DeleteRunProfileForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *RunProfilesV1DeleteRunProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunProfilesV1DeleteRunProfileNotFound creates a RunProfilesV1DeleteRunProfileNotFound with default headers values
func NewRunProfilesV1DeleteRunProfileNotFound() *RunProfilesV1DeleteRunProfileNotFound {
	return &RunProfilesV1DeleteRunProfileNotFound{}
}

/*RunProfilesV1DeleteRunProfileNotFound handles this case with default header values.

Resource does not exist.
*/
type RunProfilesV1DeleteRunProfileNotFound struct {
	Payload interface{}
}

func (o *RunProfilesV1DeleteRunProfileNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/run_profiles/{uuid}][%d] runProfilesV1DeleteRunProfileNotFound  %+v", 404, o.Payload)
}

func (o *RunProfilesV1DeleteRunProfileNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *RunProfilesV1DeleteRunProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunProfilesV1DeleteRunProfileDefault creates a RunProfilesV1DeleteRunProfileDefault with default headers values
func NewRunProfilesV1DeleteRunProfileDefault(code int) *RunProfilesV1DeleteRunProfileDefault {
	return &RunProfilesV1DeleteRunProfileDefault{
		_statusCode: code,
	}
}

/*RunProfilesV1DeleteRunProfileDefault handles this case with default header values.

An unexpected error response
*/
type RunProfilesV1DeleteRunProfileDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the run profiles v1 delete run profile default response
func (o *RunProfilesV1DeleteRunProfileDefault) Code() int {
	return o._statusCode
}

func (o *RunProfilesV1DeleteRunProfileDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/run_profiles/{uuid}][%d] RunProfilesV1_DeleteRunProfile default  %+v", o._statusCode, o.Payload)
}

func (o *RunProfilesV1DeleteRunProfileDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *RunProfilesV1DeleteRunProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}