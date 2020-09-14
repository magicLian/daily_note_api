package util

import (
	"github.com/gin-gonic/gin"
	"ml_daily_record/pkg/models"
	"strconv"
)

func GetPagation(c *gin.Context) *models.Pagination {
	page, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if limit <= 0 {
		limit = 20
	}

	pagation := &models.Pagination{
		PageNum:  page,
		Offset:   (page - 1) * limit,
		PageSize: limit,
	}
	return pagation
}
