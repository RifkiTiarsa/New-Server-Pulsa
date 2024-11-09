package controller

import (
	"server-pulsa/config"
	"server-pulsa/entity"
	"server-pulsa/shared/common"
	"server-pulsa/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type memberController struct {
	uc usecase.MemberUsecase
	rg *gin.RouterGroup
}

func (m *memberController) createHandler(c *gin.Context) {
	var payload entity.Member

	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, 400, err.Error())
		return
	}

	member, err := m.uc.Create(&payload)
	if err != nil {
		common.SendErrorResponse(c, 500, err.Error())
		return
	}

	common.SendCreateResponse(c, member)
}

func (m *memberController) getHandlerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.SendErrorResponse(c, 400, "Invalid ID")
		return
	}

	member, err := m.uc.FindByID(id)
	if err != nil {
		common.SendErrorResponse(c, 404, "Member not found")
		return
	}

	common.SendSingleResponse(c, member, "OK")
}

func (m *memberController) getAllHandler(c *gin.Context) {
	members, err := m.uc.FindAll()
	if err != nil {
		common.SendErrorResponse(c, 500, err.Error())
		return
	}

	common.SendSingleResponse(c, members, "OK")
}

func (m *memberController) updateHandler(c *gin.Context) {
	var payload entity.Member

	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, 400, err.Error())
		return
	}

	member, err := m.uc.Update(&payload)
	if err != nil {
		common.SendErrorResponse(c, 500, err.Error())
		return
	}

	common.SendSingleResponse(c, member, "Member updated successfully")
}

func (m *memberController) deleteHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.SendErrorResponse(c, 400, "Invalid ID")
		return
	}

	err = m.uc.Delete(id)
	if err != nil {
		common.SendErrorResponse(c, 500, err.Error())
		return
	}

	common.SendSingleResponse(c, nil, "Member deleted successfully")
}

func (m *memberController) Routes() {
	m.rg.POST(config.PostMember, m.createHandler)
	m.rg.GET(config.GetMember, m.getHandlerById)
	m.rg.GET(config.GetMemberList, m.getAllHandler)
	m.rg.PUT(config.PutMember, m.updateHandler)
	m.rg.DELETE(config.DeleteMember, m.deleteHandler)
}

func NewMemberController(uc usecase.MemberUsecase, rg *gin.RouterGroup) *memberController {
	return &memberController{uc: uc, rg: rg}
}
