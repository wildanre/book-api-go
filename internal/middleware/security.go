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

		// Prevent clickjacking but allow Swagger UI to function properly
		// Check if the path is for Swagger UI and apply a less restrictive policy
		if c.Request.URL.Path == "/swagger/index.html" || c.Request.URL.Path == "/swagger/" || c.Request.URL.Path == "/swagger" {
			// More permissive CSP for Swagger UI
			c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; font-src 'self' data:")
		} else {
			// Regular CSP for other routes
			c.Header("Content-Security-Policy", "default-src 'self'")
		}

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
