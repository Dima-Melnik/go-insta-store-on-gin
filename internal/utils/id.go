package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetID(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return uint(id), err
	}

	return uint(id), err
}
