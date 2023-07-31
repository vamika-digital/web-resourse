package domain

import (
	"github.com/vamika-digital/wms-resourse/pkg/constant"
	"time"
)

type User struct {
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
