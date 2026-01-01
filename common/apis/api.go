package apis

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Api struct {
	Context *gin.Context
	DB      *gorm.DB
	Errors  error
}
