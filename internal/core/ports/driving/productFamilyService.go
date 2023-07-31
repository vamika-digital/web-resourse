package driving

import (
	"github.com/vamika-digital/wms-lib/errs"
	"github.com/vamika-digital/wms-resourse/internal/core/dto/request"
	"github.com/vamika-digital/wms-resourse/internal/core/dto/response"
)

type ProductFamilyService interface {
	Save(request request.NewProductFamilyRequest) (*response.NewProductFamilyResponse, *errs.AppError)
	GetAllProductFamilies(status string) ([]response.ProductFamilyResponse, *errs.AppError)
	GetProductFamily(id string) (*response.ProductFamilyResponse, *errs.AppError)
}
