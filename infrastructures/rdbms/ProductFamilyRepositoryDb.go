package rdbms

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/vamika-digital/wms-lib/errs"
	"github.com/vamika-digital/wms-lib/logger"
	"github.com/vamika-digital/wms-resourse/internal/core/domain"
	"github.com/vamika-digital/wms-resourse/pkg/constant"
	"strconv"
	"time"
)

type ProductFamilyRepositoryDb struct {
	client *sqlx.DB
}

func NewProductFamilyRepositoryDb(dbClient *sqlx.DB) ProductFamilyRepositoryDb {
	return ProductFamilyRepositoryDb{dbClient}
}

func (r ProductFamilyRepositoryDb) FindAll(status constant.Status) ([]domain.ProductFamily, *errs.AppError) {
	var err error
	families := make([]domain.ProductFamily, 0)

	if status.IsValid() {
		findAllSql := "select id, name, description, unit, status, created_at, updated_at from product_families"
		err = r.client.Select(&families, findAllSql)
	} else {
		findAllSql := "select id, name, description, unit, status, created_at, updated_at from product_families whre status = ?"
		err = r.client.Select(&families, findAllSql, status)
	}
	if err != nil {
		logger.Error("Error while querying product_families table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return families, nil
}

func (r ProductFamilyRepositoryDb) FindById(id string) (*domain.ProductFamily, *errs.AppError) {
	familySql := "select id, name, description, unit, status, created_at, updated_at from product_families where id = ?"
	var family domain.ProductFamily

	err := r.client.Get(&family, familySql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("ProductFamily not found")
		} else {
			logger.Error("Error while scanning product family " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &family, nil
}

func (r ProductFamilyRepositoryDb) Save(family domain.ProductFamily) (*domain.ProductFamily, *errs.AppError) {
	sqlInsert := "INSERT INTO product_families (name, description, unit, status, created_at, updated_at) values (?, ?, ?, ?, ?, ?)"

	result, err := r.client.Exec(sqlInsert, family.Name, family.Description, family.Unit, family.Status, time.Now().Format(constant.GlobalDBTSLayout), time.Now().Format(constant.GlobalDBTSLayout))
	if err != nil {
		logger.Error("Error while creating new product family: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new product family: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	family.Id = strconv.FormatInt(id, 10)
	return &family, nil
}
