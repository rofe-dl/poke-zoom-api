package utils

import (
	"fmt"

	"gorm.io/gorm"
)

func Paginate(offset int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if offset <= 0 {
			offset = 0
		}

		// hard limit of 100
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		return db.Offset(offset).Limit(limit)
	}
}

func Sort(sortField string, sortOrder string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sortField == "" {
			return db
		}

		if sortOrder == "" {
			sortOrder = "asc"
		}

		return db.Order(sortField + " " + sortOrder)
	}
}

func FilterByValue[T comparable](field string, value *T) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value != nil {
			return db.Where(fmt.Sprintf("%v = ?", field), *value)
		}

		return db
	}
}

func FilterByArray[T comparable](field string, value []T) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(value) != 0 {
			return db.Where(fmt.Sprintf("%v IN ?", field), value)
		}

		return db
	}
}

func FilterByRange[T comparable](field string, min *T, max *T) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if min != nil {
			db = db.Where(fmt.Sprintf("%v >= ?", field), *min)
		}

		if max != nil {
			db = db.Where(fmt.Sprintf("%v <= ?", field), *max)
		}

		return db
	}
}

type SearchQueryParams struct {
	Search string `query:"search" example:"john" required:"false"`
}

type PaginationQueryParams struct {
	Offset    int    `query:"offset" example:"1" min:"0" required:"false"`
	Limit     int    `query:"limit" example:"10" min:"1" required:"false"`
	SortOrder string `query:"sort_order" enum:"asc,desc" example:"asc" default:"asc"`
}
