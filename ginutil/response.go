package ginutil

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/saltbo/gopkg/httputil"
)

// JSON
func JSON(c *gin.Context) {
	c.JSON(http.StatusOK, httputil.NewJSONResponse(nil))
}

// JSONError
func JSONError(c *gin.Context, status int, err error) {
	c.AbortWithStatusJSON(status, httputil.JSONResponse{
		Code: status,
		Msg:  err.Error(),
	})
	c.Error(err)
}

// JSONData
func JSONData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, httputil.NewJSONResponse(data))
}

// JSONList
func JSONList(c *gin.Context, list interface{}, total int64) {
	c.JSON(http.StatusOK, httputil.NewJSONResponse(gin.H{
		"list":  list,
		"total": total,
	}))
}

// BadRequest
func JSONBadRequest(c *gin.Context, err error) {
	JSONError(c, http.StatusBadRequest, err)
}

// Unauthorized
func JSONUnauthorized(c *gin.Context, err error) {
	JSONError(c, http.StatusUnauthorized, err)
}

// Forbidden
func JSONForbidden(c *gin.Context, err error) {
	JSONError(c, http.StatusForbidden, err)
}

// ServerError
func JSONServerError(c *gin.Context, err error) {
	JSONError(c, http.StatusInternalServerError, err)
}