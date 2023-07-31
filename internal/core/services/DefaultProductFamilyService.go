package services

import (
	"github.com/vamika-digital/wms-lib/errs"
	"github.com/vamika-digital/wms-resourse/internal/core/domain"
	"github.com/vamika-digital/wms-resourse/internal/core/dto/request"
	"github.com/vamika-digital/wms-resourse/internal/core/dto/response"
	"github.com/vamika-digital/wms-resourse/internal/core/ports/driven"
	"github.com/vamika-digital/wms-resourse/pkg/constant"
)

type DefaultProductFamilyService struct {
	repo driven.ProductFamilyRepository
}

func NewProductFamilyService(repository driven.ProductFamilyRepository) DefaultProductFamilyService {
	return DefaultProductFamilyService{repository}
}

func (s DefaultProductFamilyService) GetAllProductFamilies(status string) ([]response.ProductFamilyResponse, *errs.AppError) {
	productStatus := constant.StatusFromString(status)
	productFamilies, err := s.repo.FindAll(productStatus)
	if err != nil {
		return nil, err
	}
	response := make([]response.ProductFamilyResponse, 0)
	for _, p := range productFamilies {
		response = append(response, p.ToFullResponseDto())
	}
	return response, err
}

func (s DefaultProductFamilyService) GetProductFamily(id string) (*response.ProductFamilyResponse, *errs.AppError) {
	p, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := p.ToFullResponseDto()
	return &response, nil
}

func (s DefaultProductFamilyService) Save(request request.NewProductFamilyRequest) (*response.NewProductFamilyResponse, *errs.AppError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	productFamily := domain.NewProductFamily(request.Name, request.Description, request.Unit)
	if newProductFamily, err := s.repo.Save(productFamily); err != nil {
		return nil, err
	} else {
		response := newProductFamily.ToNewResponseDto()
		return &response, nil
	}
}
