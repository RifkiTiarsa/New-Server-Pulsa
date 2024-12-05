package common

import (
	"fmt"
	"net/http"
	"server-pulsa/entity"
	"server-pulsa/shared/model"

	"github.com/gin-gonic/gin"
)

func SendErrorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, &model.Status{
		Code:    code,
		Message: message,
	})
}

func SendCreateResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusCreated,
			Message: "Created",
		},
		Data: data,
	})
}

func SendSingleResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

func CheckBalanceMemberResponse(c *gin.Context, data entity.Member) {
	c.JSON(http.StatusOK, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: fmt.Sprintf("YTH %s %s Saldo anda: %.2f @%s", data.Name, data.MemberID, data.Balance, data.UpdatedAt),
	})
}
