///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Function function
// swagger:model Function

type Function struct {

	// code
	// Required: true
	Code *string `json:"code"`

	// created time
	CreatedTime int64 `json:"createdTime,omitempty"`

	// id
	ID strfmt.UUID `json:"id,omitempty"`

	// image
	// Required: true
	Image *string `json:"image"`

	// main
	Main *string `json:"main,omitempty"`

	// modified time
	ModifiedTime int64 `json:"modifiedTime,omitempty"`

	// name
	// Required: true
	// Pattern: ^[\w\d\-]+$
	Name *string `json:"name"`

	// schema
	Schema *Schema `json:"schema,omitempty"`

	// secrets
	Secrets []string `json:"secrets"`

	// status
	Status Status `json:"status,omitempty"`

	// tags
	Tags FunctionTags `json:"tags"`
}

/* polymorph Function code false */

/* polymorph Function createdTime false */

/* polymorph Function id false */

/* polymorph Function image false */

/* polymorph Function main false */

/* polymorph Function modifiedTime false */

/* polymorph Function name false */

/* polymorph Function schema false */

/* polymorph Function secrets false */

/* polymorph Function status false */

/* polymorph Function tags false */

// Validate validates this function
func (m *Function) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecrets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Function) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *Function) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *Function) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Function) validateSchema(formats strfmt.Registry) error {

	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {

		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

func (m *Function) validateSecrets(formats strfmt.Registry) error {

	if swag.IsZero(m.Secrets) { // not required
		return nil
	}

	return nil
}

func (m *Function) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Function) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Function) UnmarshalBinary(b []byte) error {
	var res Function
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
