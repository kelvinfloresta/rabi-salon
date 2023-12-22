package gorm_adapter

import (
	"rabi-salon/frameworks/database"

	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, paginate database.PaginateInput) {
	offset := paginate.Page * paginate.PageSize
	db.Limit(paginate.PageSize).Offset(offset)
}
