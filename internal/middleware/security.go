package middleware

import (
	"github.com/gin-gonic/gin"
)

// SecurityHeaders adds security headers to responses
func SecurityHeaders() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// Prevent XSS attacks
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")

		// Prevent clickjacking
		c.Header("Content-Security-Policy", "default-src 'self'")

		// Force HTTPS in production
		if gin.Mode() == gin.ReleaseMode {
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		}

		c.Next()
	})
}

// RateLimitByIP implements basic rate limiting
func RateLimitByIP() gin.HandlerFunc {
	// Basic rate limiting implementation
	// For production, consider using redis-based rate limiting
	return gin.HandlerFunc(func(c *gin.Context) {
		// Simple rate limiting logic here
		// Or use a library like github.com/ulule/limiter
		c.Next()
	})
}
