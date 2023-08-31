package handler

import (
	"avito-app"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get User Segments By ID
// @Tags User
// @Description Returns segments the user with given UserID is in
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

// @Summary View User's Actions
// @Tags User
// @Description Get Actions the User was involved in
// @ID get-actions
// @Accept  json
// @Produce  json
// @Param input body avito.UserGetActions true "User ID and Date (Month and Year)"
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
// @Description Adds user with a given UserID to specified segment(s)
// @ID add-user-to-segments
// @Accept  json
// @Produce  json
// @Param input body avito.UserAddToSegment true "UserID and segment name(s)"
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
// @Description Removes user with a given UserID from specified segment(s)
// @ID remove-user-from-segments
// @Accept  json
// @Produce  json
// @Param input body avito.UserRemoveFromSegment true "UserID and segment name(s)"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/removeFromSegment [post]
func (h *Handler) removeUserFromSegment(c *gin.Context) {
	var input avito.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.User.RemoveUserFromSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
