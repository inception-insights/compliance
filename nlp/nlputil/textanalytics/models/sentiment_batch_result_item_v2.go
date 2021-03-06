package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// SentimentBatchResultItemV2 sentiment batch result item v2
// swagger:model SentimentBatchResultItemV2
type SentimentBatchResultItemV2 struct {

	// Unique document identifier.
	ID string `json:"id,omitempty"`

	// A decimal number between 0 and 1 denoting the sentiment of the document.
	//             A score above 0.7 usually refers to a positive document while a score below 0.3 normally has a negative connotation.
	//             Mid values refer to neutral text.
	Score float64 `json:"score,omitempty"`
}

// Validate validates this sentiment batch result item v2
func (m *SentimentBatchResultItemV2) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
