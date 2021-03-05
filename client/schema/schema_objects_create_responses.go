//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// SchemaObjectsCreateReader is a Reader for the SchemaObjectsCreate structure.
type SchemaObjectsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaObjectsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchemaObjectsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSchemaObjectsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSchemaObjectsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSchemaObjectsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSchemaObjectsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSchemaObjectsCreateOK creates a SchemaObjectsCreateOK with default headers values
func NewSchemaObjectsCreateOK() *SchemaObjectsCreateOK {
	return &SchemaObjectsCreateOK{}
}

/*SchemaObjectsCreateOK handles this case with default header values.

Added the new Object class to the schema.
*/
type SchemaObjectsCreateOK struct {
	Payload *models.Class
}

func (o *SchemaObjectsCreateOK) Error() string {
	return fmt.Sprintf("[POST /schema][%d] schemaObjectsCreateOK  %+v", 200, o.Payload)
}

func (o *SchemaObjectsCreateOK) GetPayload() *models.Class {
	return o.Payload
}

func (o *SchemaObjectsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Class)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsCreateUnauthorized creates a SchemaObjectsCreateUnauthorized with default headers values
func NewSchemaObjectsCreateUnauthorized() *SchemaObjectsCreateUnauthorized {
	return &SchemaObjectsCreateUnauthorized{}
}

/*SchemaObjectsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type SchemaObjectsCreateUnauthorized struct {
}

func (o *SchemaObjectsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /schema][%d] schemaObjectsCreateUnauthorized ", 401)
}

func (o *SchemaObjectsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaObjectsCreateForbidden creates a SchemaObjectsCreateForbidden with default headers values
func NewSchemaObjectsCreateForbidden() *SchemaObjectsCreateForbidden {
	return &SchemaObjectsCreateForbidden{}
}

/*SchemaObjectsCreateForbidden handles this case with default header values.

Forbidden
*/
type SchemaObjectsCreateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *SchemaObjectsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /schema][%d] schemaObjectsCreateForbidden  %+v", 403, o.Payload)
}

func (o *SchemaObjectsCreateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsCreateUnprocessableEntity creates a SchemaObjectsCreateUnprocessableEntity with default headers values
func NewSchemaObjectsCreateUnprocessableEntity() *SchemaObjectsCreateUnprocessableEntity {
	return &SchemaObjectsCreateUnprocessableEntity{}
}

/*SchemaObjectsCreateUnprocessableEntity handles this case with default header values.

Invalid Object class
*/
type SchemaObjectsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *SchemaObjectsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /schema][%d] schemaObjectsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SchemaObjectsCreateUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsCreateInternalServerError creates a SchemaObjectsCreateInternalServerError with default headers values
func NewSchemaObjectsCreateInternalServerError() *SchemaObjectsCreateInternalServerError {
	return &SchemaObjectsCreateInternalServerError{}
}

/*SchemaObjectsCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaObjectsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SchemaObjectsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /schema][%d] schemaObjectsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaObjectsCreateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}