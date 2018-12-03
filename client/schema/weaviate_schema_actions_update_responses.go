// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateSchemaActionsUpdateReader is a Reader for the WeaviateSchemaActionsUpdate structure.
type WeaviateSchemaActionsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaActionsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaActionsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateSchemaActionsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateSchemaActionsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateSchemaActionsUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaActionsUpdateOK creates a WeaviateSchemaActionsUpdateOK with default headers values
func NewWeaviateSchemaActionsUpdateOK() *WeaviateSchemaActionsUpdateOK {
	return &WeaviateSchemaActionsUpdateOK{}
}

/*WeaviateSchemaActionsUpdateOK handles this case with default header values.

Changes applied
*/
type WeaviateSchemaActionsUpdateOK struct {
}

func (o *WeaviateSchemaActionsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /schema/actions/{className}][%d] weaviateSchemaActionsUpdateOK ", 200)
}

func (o *WeaviateSchemaActionsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaActionsUpdateUnauthorized creates a WeaviateSchemaActionsUpdateUnauthorized with default headers values
func NewWeaviateSchemaActionsUpdateUnauthorized() *WeaviateSchemaActionsUpdateUnauthorized {
	return &WeaviateSchemaActionsUpdateUnauthorized{}
}

/*WeaviateSchemaActionsUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaActionsUpdateUnauthorized struct {
}

func (o *WeaviateSchemaActionsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /schema/actions/{className}][%d] weaviateSchemaActionsUpdateUnauthorized ", 401)
}

func (o *WeaviateSchemaActionsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaActionsUpdateForbidden creates a WeaviateSchemaActionsUpdateForbidden with default headers values
func NewWeaviateSchemaActionsUpdateForbidden() *WeaviateSchemaActionsUpdateForbidden {
	return &WeaviateSchemaActionsUpdateForbidden{}
}

/*WeaviateSchemaActionsUpdateForbidden handles this case with default header values.

Could not find the Action class
*/
type WeaviateSchemaActionsUpdateForbidden struct {
}

func (o *WeaviateSchemaActionsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /schema/actions/{className}][%d] weaviateSchemaActionsUpdateForbidden ", 403)
}

func (o *WeaviateSchemaActionsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaActionsUpdateUnprocessableEntity creates a WeaviateSchemaActionsUpdateUnprocessableEntity with default headers values
func NewWeaviateSchemaActionsUpdateUnprocessableEntity() *WeaviateSchemaActionsUpdateUnprocessableEntity {
	return &WeaviateSchemaActionsUpdateUnprocessableEntity{}
}

/*WeaviateSchemaActionsUpdateUnprocessableEntity handles this case with default header values.

Invalid update
*/
type WeaviateSchemaActionsUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaActionsUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /schema/actions/{className}][%d] weaviateSchemaActionsUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateSchemaActionsUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*WeaviateSchemaActionsUpdateBody weaviate schema actions update body
swagger:model WeaviateSchemaActionsUpdateBody
*/
type WeaviateSchemaActionsUpdateBody struct {

	// keywords
	Keywords models.SemanticSchemaKeywords `json:"keywords"`

	// The new name of the Action
	NewName string `json:"newName,omitempty"`
}

// Validate validates this weaviate schema actions update body
func (o *WeaviateSchemaActionsUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKeywords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateSchemaActionsUpdateBody) validateKeywords(formats strfmt.Registry) error {

	if swag.IsZero(o.Keywords) { // not required
		return nil
	}

	if err := o.Keywords.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "keywords")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateSchemaActionsUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateSchemaActionsUpdateBody) UnmarshalBinary(b []byte) error {
	var res WeaviateSchemaActionsUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}