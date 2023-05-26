package handler

import (
	"bicycle/bicycle_api_gateway/api/http"
	"bicycle/bicycle_api_gateway/genproto/order_service"

	"github.com/gin-gonic/gin"
)

// Create Order godoc
// @ID create_order
// @Router /order [POST]
// @Summary Create Order
// @Description Create Order
// @Tags Order
// @Accept json
// @Produce json
// @Param order body order_service.Order true "CreateOrderRequest"
// @Success 200 {object} order_service.Order "Success Request"
// @Response 400 {object} order_service.Order "Bad Request"
// @Failure 500 {object} order_service.Order "Server Error"
func (h *Handler) CreateOrder(c *gin.Context) {
	var order order_service.CreateOrder

	err := c.ShouldBindJSON(&order)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	resp, err := h.services.OrderService().Create(
		c.Request.Context(),
		&order,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Get By ID Order godoc
// @ID get_order_by_id
// @Router /order/{id} [GET]
// @Summary Get By ID Order
// @Description Get By ID Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} order_service.Order "Success Request"
// @Response 400 {object} order_service.Order "Bad Request"
// @Failure 500 {object} order_service.Order "Server Error"
func (h *Handler) GetOrderById(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.OrderService().GetById(
		c.Request.Context(),
		&order_service.PrimaryKey{
			Id: id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DELETE Order godoc
// @ID delete_order
// @Router /order/{id} [DELETE]
// @Summary Delete Order
// @Description Delete Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Empty
func (h *Handler) DeleteOrder(c *gin.Context) {
	var resp http.Empty

	id := c.Param("id")

	_, err := h.services.OrderService().Delete(
		c.Request.Context(),
		&order_service.PrimaryKey{
			Id: id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
