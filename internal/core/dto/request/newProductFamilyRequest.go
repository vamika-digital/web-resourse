package request

import (
	"github.com/vamika-digital/wms-lib/errs"
	"github.com/vamika-digital/wms-resourse/pkg/constant"
)

type NewProductFamilyRequest struct {
	Name        string
	Description string
	Unit        constant.Measurement
}

func (r NewProductFamilyRequest) Validate() *errs.AppError {
	if r.Name == "" {
		return errs.NewValidationError("Name is required")
	} else if len(r.Name) < 3 || len(r.Name) > 20 {
		return errs.NewValidationError("Name length should be in between 3 and 20")
	}

	if r.Description != "" && len(r.Name) > 200 {
		return errs.NewValidationError("Description must be in limit of 200 characters")
	}

	if !r.Unit.IsValid() {
		return errs.NewValidationError("Unit type is invalid")
	}
	return nil
}
