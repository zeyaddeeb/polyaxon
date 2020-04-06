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

// RunProfilesV1PatchRunProfileReader is a Reader for the RunProfilesV1PatchRunProfile structure.
type RunProfilesV1PatchRunProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunProfilesV1PatchRunProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunProfilesV1PatchRunProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewRunProfilesV1PatchRunProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRunProfilesV1PatchRunProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunProfilesV1PatchRunProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRunProfilesV1PatchRunProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunProfilesV1PatchRunProfileOK creates a RunProfilesV1PatchRunProfileOK with default headers values
func NewRunProfilesV1PatchRunProfileOK() *RunProfilesV1PatchRunProfileOK {
	return &RunProfilesV1PatchRunProfileOK{}
}

/*RunProfilesV1PatchRunProfileOK handles this case with default header values.

A successful response.
*/
type RunProfilesV1PatchRunProfileOK struct {
	Payload *service_model.V1RunProfile
}

func (o *RunProfilesV1PatchRunProfileOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}][%d] runProfilesV1PatchRunProfileOK  %+v", 200, o.Payload)
}

func (o *RunProfilesV1PatchRunProfileOK) GetPayload() *service_model.V1RunProfile {
	return o.Payload
}

func (o *RunProfilesV1PatchRunProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1RunProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunProfilesV1PatchRunProfileNoContent creates a RunProfilesV1PatchRunProfileNoContent with default headers values
func NewRunProfilesV1PatchRunProfileNoContent() *RunProfilesV1PatchRunProfileNoContent {
	return &RunProfilesV1PatchRunProfileNoContent{}
}

/*RunProfilesV1PatchRunProfileNoContent handles this case with default header values.

No content.
*/
type RunProfilesV1PatchRunProfileNoContent struct {
	Payload interface{}
}

func (o *RunProfilesV1PatchRunProfileNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}][%d] runProfilesV1PatchRunProfileNoContent  %+v", 204, o.Payload)
}

func (o *RunProfilesV1PatchRunProfileNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *RunProfilesV1PatchRunProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunProfilesV1PatchRunProfileForbidden creates a RunProfilesV1PatchRunProfileForbidden with default headers values
func NewRunProfilesV1PatchRunProfileForbidden() *RunProfilesV1PatchRunProfileForbidden {
	return &RunProfilesV1PatchRunProfileForbidden{}
}

/*RunProfilesV1PatchRunProfileForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type RunProfilesV1PatchRunProfileForbidden struct {
	Payload interface{}
}

func (o *RunProfilesV1PatchRunProfileForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}][%d] runProfilesV1PatchRunProfileForbidden  %+v", 403, o.Payload)
}

func (o *RunProfilesV1PatchRunProfileForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *RunProfilesV1PatchRunProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunProfilesV1PatchRunProfileNotFound creates a RunProfilesV1PatchRunProfileNotFound with default headers values
func NewRunProfilesV1PatchRunProfileNotFound() *RunProfilesV1PatchRunProfileNotFound {
	return &RunProfilesV1PatchRunProfileNotFound{}
}

/*RunProfilesV1PatchRunProfileNotFound handles this case with default header values.

Resource does not exist.
*/
type RunProfilesV1PatchRunProfileNotFound struct {
	Payload interface{}
}

func (o *RunProfilesV1PatchRunProfileNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}][%d] runProfilesV1PatchRunProfileNotFound  %+v", 404, o.Payload)
}

func (o *RunProfilesV1PatchRunProfileNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *RunProfilesV1PatchRunProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunProfilesV1PatchRunProfileDefault creates a RunProfilesV1PatchRunProfileDefault with default headers values
func NewRunProfilesV1PatchRunProfileDefault(code int) *RunProfilesV1PatchRunProfileDefault {
	return &RunProfilesV1PatchRunProfileDefault{
		_statusCode: code,
	}
}

/*RunProfilesV1PatchRunProfileDefault handles this case with default header values.

An unexpected error response
*/
type RunProfilesV1PatchRunProfileDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the run profiles v1 patch run profile default response
func (o *RunProfilesV1PatchRunProfileDefault) Code() int {
	return o._statusCode
}

func (o *RunProfilesV1PatchRunProfileDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}][%d] RunProfilesV1_PatchRunProfile default  %+v", o._statusCode, o.Payload)
}

func (o *RunProfilesV1PatchRunProfileDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *RunProfilesV1PatchRunProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}