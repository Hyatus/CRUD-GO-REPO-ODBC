package middleware

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Hyatus/myapi/utils"
)


func ErrorHandler() gin.HandlerFunc { 
	return func ( c *gin.Context) {
		defer func(){
			if err := recover(); err != nil {
				utils.Log.WithFields(map[string]interface{}{
					"method": c.Request.Method,
					"path": c.Request.URL.Path, 
					"error": err,
				})
				log.Printf("Internal error %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Internal server error", 
					"code": 500, 
				})
			}
		}()
		c.Next() // Proceed to the next middleware or handler
	}
}