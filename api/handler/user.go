package handler

import (
	"bicycle/bicycle_api_gateway/api/http"
	"bicycle/bicycle_api_gateway/genproto/user_service"

	"github.com/gin-gonic/gin"
)

// Create User godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body user_service.User true "CreateUserRequest"
// @Success 200 {object} user_service.User "Success Request"
// @Response 400 {object} user_service.User "Bad Request"
// @Failure 500 {object} user_service.User "Server Error"
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

// Get By ID User godoc
// @ID get_user_by_id
// @Router /user/{id} [GET]
// @Summary Get By ID User
// @Description Get By ID User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} user_service.User "Success Request"
// @Response 400 {object} user_service.User "Bad Request"
// @Failure 500 {object} user_service.User "Server Error"
func (h *Handler) GetUserById(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.UserService().GetById(
		c.Request.Context(),
		&user_service.PrimaryKey{
			Id: id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DELETE User godoc
// @ID delete_user
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
func (h *Handler) DeleteUser(c *gin.Context) {
	var resp http.Empty

	id := c.Param("id")

	_, err := h.services.UserService().Delete(
		c.Request.Context(),
		&user_service.PrimaryKey{
			Id: id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
