package middleware

import "github.com/gin-gonic/gin"

// CORS will handle the CORS middleware
func CORS() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}
		context.Next()
	}
}
