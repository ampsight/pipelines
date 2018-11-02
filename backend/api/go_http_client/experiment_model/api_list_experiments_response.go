// Copyright 2018 Google LLC
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

package experiment_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APIListExperimentsResponse api list experiments response
// swagger:model apiListExperimentsResponse
type APIListExperimentsResponse struct {

	// experiments
	Experiments []*APIExperiment `json:"experiments"`

	// next page token
	NextPageToken string `json:"next_page_token,omitempty"`
}

// Validate validates this api list experiments response
func (m *APIListExperimentsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExperiments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIListExperimentsResponse) validateExperiments(formats strfmt.Registry) error {

	if swag.IsZero(m.Experiments) { // not required
		return nil
	}

	for i := 0; i < len(m.Experiments); i++ {
		if swag.IsZero(m.Experiments[i]) { // not required
			continue
		}

		if m.Experiments[i] != nil {
			if err := m.Experiments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("experiments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIListExperimentsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIListExperimentsResponse) UnmarshalBinary(b []byte) error {
	var res APIListExperimentsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}