package apis

import (
	models2 "TaoCoupon/app/models"
	"TaoCoupon/app/server/models"
	"TaoCoupon/app/server/models/request"
	"TaoCoupon/common/apis"
	"TaoCoupon/common/database"
	"TaoCoupon/common/logger"
	page2 "TaoCoupon/common/page"
	"TaoCoupon/common/redis"
	"TaoCoupon/common/snowflake"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type CouponTemplate struct {
	apis.Api
}

func (a CouponTemplate) CreateCouponTemplate(c *gin.Context) {
	var req request.CouponTemplateInsertReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Error("request param fail")
		panic(fmt.Errorf("request param error, %v", err))
	}

	coupon := models.CouponTemplate{
		Name:           req.Name,
		Source:         req.Source,
		ShopNumber:     req.ShopNumber,
		Target:         req.Target,
		Goods:          req.Goods,
		Type:           req.Type,
		ValidStartTime: req.ValidStartTime,
		ValidEndTime:   req.ValidEndTime,
		Stock:          req.Stock,
		ReceiveRule:    req.ReceiveRule,
		ConsumeRule:    req.ConsumeRule,
		Status:         0,                   // 默认：生效中
		ModelTime:      models2.ModelTime{}, // CreatedAt / UpdatedAt 自动写入
	}

	coupon.Id = snowflake.IdGen.NextID()

	if err := database.DB.Create(&coupon).Error; err != nil {
		logger.Logger.Error("create coupon fail: ", err)
		c.JSON(500, gin.H{"msg": "insert failed", "err": err.Error()})
		return
	}
	key := fmt.Sprintf("coupon_template:%d", coupon.Id)

	body, err := json.Marshal(coupon)
	if err != nil {
		logger.Logger.Error("marshal coupon fail: ", err)
	}

	// 设置 key + 过期时间
	expire := time.Until(coupon.ValidEndTime)
	if expire <= 0 {
		expire = 1 * time.Second
	}

	redis.RedisClient.Set(context.Background(), key, body, expire)

}

func (a CouponTemplate) PageQueryCouponTemplate(c *gin.Context) {
	var req request.CouponTemplateQueryReq
	var page page2.PageReq

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"msg": "参数错误"})
		return
	}

	_ = c.ShouldBindQuery(&page)
	fmt.Println(page, req)
	page.Paginate(page.Page, page.PageSize)
	db := database.DB.Model(&models.CouponTemplate{})

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Target != 0 {
		db = db.Where("target = ?", req.Target)
	}
	if req.Goods != 0 {
		db = db.Where("goods = ?", req.Goods)
	}
	if req.Type != 0 {
		db = db.Where("type = ?", req.Type)
	}

	var total int64
	db.Count(&total)

	var list []models.CouponTemplate
	db.Limit(page.PageSize).
		Offset((page.Page - 1) * page.PageSize).
		Order("id DESC").
		Find(&list)

	c.JSON(200, gin.H{
		"total":     total,
		"page":      page.Page,
		"page_size": page.PageSize,
		"list":      list,
	})
}

func (a CouponTemplate) FindCouponTemplate(c *gin.Context) {

}

func (a CouponTemplate) IncreaseNumberCouponTemplate(c *gin.Context) {

}

func (a CouponTemplate) TerminateCouponTemplate(c *gin.Context) {

}
