package domain

import (
	"github.com/vamika-digital/wms-resourse/pkg/constant"
	"time"
)

type ProductType int8

const (
	RawMaterial  ProductType = 1
	Manufactured ProductType = 2
)

type Product struct {
	Id          string
	Name        string
	Description string
	Unit        constant.Measurement
	Type        ProductType
	Family      ProductFamily
	Components  []Product
	Status      constant.Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
