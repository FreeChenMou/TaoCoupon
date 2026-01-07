package page

import "gorm.io/gorm"

type PageReq struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

func (p PageReq) Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		if pageSize > 100 {
			pageSize = 100
		}
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}
