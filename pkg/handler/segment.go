package handler

import (
	"avito-app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a Segment
// @Tags Segment
// @Description Create a Segment
// @ID create-segment
// @Accept  json
// @Produce  json
// @Param input body avito.Segment true "Segment's name"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segment/create [post]
func (h *Handler) createSegment(c *gin.Context) {
	var input avito.Segment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Segment.CreateSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Delete a Segment
// @Tags Segment
// @Description Delete a Segment
// @ID delete-segment
// @Accept  json
// @Produce  json
// @Param input body avito.Segment true "Segment's name"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segment/delete [post]
func (h *Handler) deleteSegment(c *gin.Context) {
	var input avito.Segment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Segment.DeleteSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"deleted": id,
	})
}
