package driven

import (
	"github.com/vamika-digital/wms-lib/errs"
	"github.com/vamika-digital/wms-resourse/internal/core/domain"
	"github.com/vamika-digital/wms-resourse/pkg/constant"
)

type ProductFamilyRepository interface {
	Save(family domain.ProductFamily) (domain.ProductFamily, *errs.AppError)
	FindAll(status constant.Status) ([]domain.ProductFamily, *errs.AppError)
	FindById(id string) (*domain.ProductFamily, *errs.AppError)
}
