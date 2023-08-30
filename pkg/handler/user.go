package handler

import (
	"avito-app"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get User Segments By ID
// @Tags User
// @Description Get Segments the User is in
// @ID get-segments
// @Accept  json
// @Produce  json
// @Param id path integer true "User's ID"
// @Success 200 {array} string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/getSegments/{id} [get]
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

// @Summary Get User Actions By ID and Date
// @Tags User
// @Description Get Actions the User was involved in
// @ID get-actions
// @Accept  json
// @Produce  json
// @Param input body avito.User true "User ID and Date"
// @Success 200 {array} string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/getActions/ [post]
func (h *Handler) getUserActions(c *gin.Context) {
	var input avito.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	slice, err := h.services.GetUserActions(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, slice)
}

// @Summary Add User to Segment(s)
// @Tags User
// @Description Add Segments to User's list
// @ID add-user-to-segments
// @Accept  json
// @Produce  json
// @Param input body avito.User true "User ID and segments"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/addToSegment [post]
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

// @Summary Remove User from Segment(s)
// @Tags User
// @Description Remove Segments from User's list
// @ID remove-user-from-segments
// @Accept  json
// @Produce  json
// @Param input body avito.User true "User ID and segments"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/deleteFromSegment [post]
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
