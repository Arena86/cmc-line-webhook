package libs

import "github.com/gin-gonic/gin"

// RecoverError func
func RecoverError(c *gin.Context) {
	if r := recover(); r != nil {
		responseData := gin.H{
			"status": 500,
			"msg":    r,
		}
		c.JSON(500, responseData)
		return
	}
}

// APIResponseData func
func APIResponseData(c *gin.Context, status int, responseData gin.H) {
	responseType := c.Request.Header.Get("ResponseType")
	if responseType == "application/xml" {
		c.XML(status, responseData)
	} else {
		c.JSON(status, responseData)
	}
}
