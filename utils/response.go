package utils

import "github.com/gin-gonic/gin"

type ErrorResponse struct { 
	Message string `json:"message"`
	Code    int    `json:"code"`
}


func RespondWithError(c *gin.Context, code int, message string){
	Log.WithFields(map[string]interface{}{
		"status": code, 
		"message": message,
		"path": c.Request.URL.Path, 
		"method": c.Request.Method,
	}).Error("API Error Response")
	
	c.AbortWithStatusJSON(code, ErrorResponse{
		Message: message,
		Code:    code,
	})
}

func RespondWithSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{ 
		"data": data,
	})
 }