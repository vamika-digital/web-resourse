package domain

import (
	"github.com/vamika-digital/wms-resourse/internal/core/dto/response"
	"github.com/vamika-digital/wms-resourse/pkg/constant"
	"time"
)

type ProductFamily struct {
	Id          string
	Name        string
	Description string
	Unit        constant.Measurement
	Status      constant.Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProductFamily(name string, description string, unit constant.Measurement) ProductFamily {
	return ProductFamily{
		Name:        name,
		Description: description,
		Unit:        unit,
		Status:      constant.Enable,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (r ProductFamily) ToFullResponseDto() response.ProductFamilyResponse {
	return response.ProductFamilyResponse{
		Id:          r.Id,
		Name:        r.Name,
		Description: r.Description,
		Unit:        r.Unit.AsText(),
		Status:      r.Status.AsText(),
		CreatedAt:   r.CreatedAt.Format(constant.GlobalDBTSLayout),
		UpdatedAt:   r.UpdatedAt.Format(constant.GlobalDBTSLayout),
	}
}

func (r ProductFamily) ToNewResponseDto() response.NewProductFamilyResponse {
	return response.NewProductFamilyResponse{
		Id: r.Id,
	}
}
