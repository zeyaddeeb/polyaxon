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

package project_searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewProjectSearchesV1ListProjectSearchNamesParams creates a new ProjectSearchesV1ListProjectSearchNamesParams object
// with the default values initialized.
func NewProjectSearchesV1ListProjectSearchNamesParams() *ProjectSearchesV1ListProjectSearchNamesParams {
	var ()
	return &ProjectSearchesV1ListProjectSearchNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectSearchesV1ListProjectSearchNamesParamsWithTimeout creates a new ProjectSearchesV1ListProjectSearchNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectSearchesV1ListProjectSearchNamesParamsWithTimeout(timeout time.Duration) *ProjectSearchesV1ListProjectSearchNamesParams {
	var ()
	return &ProjectSearchesV1ListProjectSearchNamesParams{

		timeout: timeout,
	}
}

// NewProjectSearchesV1ListProjectSearchNamesParamsWithContext creates a new ProjectSearchesV1ListProjectSearchNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectSearchesV1ListProjectSearchNamesParamsWithContext(ctx context.Context) *ProjectSearchesV1ListProjectSearchNamesParams {
	var ()
	return &ProjectSearchesV1ListProjectSearchNamesParams{

		Context: ctx,
	}
}

// NewProjectSearchesV1ListProjectSearchNamesParamsWithHTTPClient creates a new ProjectSearchesV1ListProjectSearchNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectSearchesV1ListProjectSearchNamesParamsWithHTTPClient(client *http.Client) *ProjectSearchesV1ListProjectSearchNamesParams {
	var ()
	return &ProjectSearchesV1ListProjectSearchNamesParams{
		HTTPClient: client,
	}
}

/*ProjectSearchesV1ListProjectSearchNamesParams contains all the parameters to send to the API endpoint
for the project searches v1 list project search names operation typically these are written to a http.Request
*/
type ProjectSearchesV1ListProjectSearchNamesParams struct {

	/*Limit
	  Limit size.

	*/
	Limit *int32
	/*Offset
	  Pagination offset.

	*/
	Offset *int32
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project under namesapce

	*/
	Project string
	/*Query
	  Query filter the search search.

	*/
	Query *string
	/*Sort
	  Sort to order the search.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) WithTimeout(timeout time.Duration) *ProjectSearchesV1ListProjectSearchNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) WithContext(ctx context.Context) *ProjectSearchesV1ListProjectSearchNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) WithHTTPClient(client *http.Client) *ProjectSearchesV1ListProjectSearchNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) WithLimit(limit *int32) *ProjectSearchesV1ListProjectSearchNamesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) WithOffset(offset *int32) *ProjectSearchesV1ListProjectSearchNamesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOwner adds the owner to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) WithOwner(owner string) *ProjectSearchesV1ListProjectSearchNamesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) WithProject(project string) *ProjectSearchesV1ListProjectSearchNamesParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) SetProject(project string) {
	o.Project = project
}

// WithQuery adds the query to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) WithQuery(query *string) *ProjectSearchesV1ListProjectSearchNamesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) WithSort(sort *string) *ProjectSearchesV1ListProjectSearchNamesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the project searches v1 list project search names params
func (o *ProjectSearchesV1ListProjectSearchNamesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectSearchesV1ListProjectSearchNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}