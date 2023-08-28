package handler

import (
	"avito-app"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userSegmentsSlice struct {
	segments []avito.Segment `json:"segments"`
}

func (h *Handler) getUserSegments(c *gin.Context) {
	usr_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	slice, err := h.services.GetUserSegments(usr_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, slice)
}

func (h *Handler) addUserToSegment(c *gin.Context) {
	var input avito.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.User.AddUserToSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteUserFromSegment(c *gin.Context) {
	var input avito.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.User.DeleteUserFromSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
