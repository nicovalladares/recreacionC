package handler

import (
	"main/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type HandleLME interface {
	Ping(c *gin.Context)

}
type handleLME struct {
	service services.ServiceLME
}

func NuevoHandle(db *gorm.DB) HandleLME {
	return handleLME{services.NewServiceLME(db)}
}


func (h handleLME) Ping(ctx *gin.Context) {
	id_li := "1-11111111"

	dto, err := h.service.GetLME(ctx.Request.Context(), id_li)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200,dto)
}