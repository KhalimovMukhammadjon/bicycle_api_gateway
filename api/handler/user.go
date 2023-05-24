package handler

import (
	"bicycle/bicycle_api_gateway/api/http"
	"bicycle/bicycle_api_gateway/genproto/user_service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user user_service.CreateUserRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
	}

	resp, err := h.services.UserService().Create(
		c.Request.Context(),
		&user,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
	}

	h.handleResponse(c, http.OK, resp)
}

func (h *Handler) GetById(c *gin.Context) {

}

// func (h *Handler)
