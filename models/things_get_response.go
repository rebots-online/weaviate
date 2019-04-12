/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ThingsGetResponse things get response
// swagger:model ThingsGetResponse
type ThingsGetResponse struct {
	Thing

	// result
	Result *ThingsGetResponseAO1Result `json:"result,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ThingsGetResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Thing
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Thing = aO0

	// AO1
	var dataAO1 struct {
		Result *ThingsGetResponseAO1Result `json:"result,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Result = dataAO1.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ThingsGetResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Thing)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Result *ThingsGetResponseAO1Result `json:"result,omitempty"`
	}

	dataAO1.Result = m.Result

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this things get response
func (m *ThingsGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Thing
	if err := m.Thing.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThingsGetResponse) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThingsGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThingsGetResponse) UnmarshalBinary(b []byte) error {
	var res ThingsGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ThingsGetResponseAO1Result Results for this specific Thing.
// swagger:model ThingsGetResponseAO1Result
type ThingsGetResponseAO1Result struct {

	// errors
	Errors *ErrorResponse `json:"errors,omitempty"`

	// status
	// Enum: [SUCCESS PENDING FAILED]
	Status *string `json:"status,omitempty"`
}

// Validate validates this things get response a o1 result
func (m *ThingsGetResponseAO1Result) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThingsGetResponseAO1Result) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	if m.Errors != nil {
		if err := m.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

var thingsGetResponseAO1ResultTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","PENDING","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		thingsGetResponseAO1ResultTypeStatusPropEnum = append(thingsGetResponseAO1ResultTypeStatusPropEnum, v)
	}
}

const (

	// ThingsGetResponseAO1ResultStatusSUCCESS captures enum value "SUCCESS"
	ThingsGetResponseAO1ResultStatusSUCCESS string = "SUCCESS"

	// ThingsGetResponseAO1ResultStatusPENDING captures enum value "PENDING"
	ThingsGetResponseAO1ResultStatusPENDING string = "PENDING"

	// ThingsGetResponseAO1ResultStatusFAILED captures enum value "FAILED"
	ThingsGetResponseAO1ResultStatusFAILED string = "FAILED"
)

// prop value enum
func (m *ThingsGetResponseAO1Result) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, thingsGetResponseAO1ResultTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ThingsGetResponseAO1Result) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("result"+"."+"status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThingsGetResponseAO1Result) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThingsGetResponseAO1Result) UnmarshalBinary(b []byte) error {
	var res ThingsGetResponseAO1Result
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
