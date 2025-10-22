package controllers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/TiginSawala-del/go-crud.git/initializers"
	"github.com/gin-gonic/gin"
)

// HealthCheck provides basic health status
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "Server is running",
		"time":    time.Now().Format(time.RFC3339),
	})
}

// HealthCheckDetailed provides detailed health information
func HealthCheckDetailed(c *gin.Context) {
	// Check database connection
	dbStatus := "connected"
	dbError := ""

	if initializers.DB != nil {
		sqlDB, err := initializers.DB.DB()
		if err != nil {
			dbStatus = "error"
			dbError = err.Error()
		} else {
			err = sqlDB.Ping()
			if err != nil {
				dbStatus = "disconnected"
				dbError = err.Error()
			} else {
				// Get database stats
				stats := sqlDB.Stats()
				c.JSON(http.StatusOK, gin.H{
					"status":  "healthy",
					"message": "All systems operational",
					"time":    time.Now().Format(time.RFC3339),
					"database": gin.H{
						"status":              dbStatus,
						"max_open_conns":      stats.MaxOpenConnections,
						"open_conns":          stats.OpenConnections,
						"in_use":              stats.InUse,
						"idle":                stats.Idle,
						"wait_count":          stats.WaitCount,
						"wait_duration":       stats.WaitDuration.String(),
						"max_idle_closed":     stats.MaxIdleClosed,
						"max_lifetime_closed": stats.MaxLifetimeClosed,
					},
					"system": gin.H{
						"goroutines":   runtime.NumGoroutine(),
						"go_version":   runtime.Version(),
						"memory_alloc": formatBytes(getMemStats().Alloc),
						"memory_total": formatBytes(getMemStats().TotalAlloc),
						"memory_sys":   formatBytes(getMemStats().Sys),
						"num_gc":       getMemStats().NumGC,
					},
				})
				return
			}
		}
	} else {
		dbStatus = "not initialized"
		dbError = "Database connection is nil"
	}

	// If we reach here, something is wrong
	c.JSON(http.StatusServiceUnavailable, gin.H{
		"status":  "unhealthy",
		"message": "Service unavailable",
		"time":    time.Now().Format(time.RFC3339),
		"database": gin.H{
			"status": dbStatus,
			"error":  dbError,
		},
		"system": gin.H{
			"goroutines": runtime.NumGoroutine(),
			"go_version": runtime.Version(),
		},
	})
}

// HealthCheckReadiness checks if the service is ready to accept traffic
func HealthCheckReadiness(c *gin.Context) {
	// Check database
	if initializers.DB == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "not ready",
			"message": "Database not initialized",
			"ready":   false,
		})
		return
	}

	sqlDB, err := initializers.DB.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "not ready",
			"message": "Database connection error",
			"error":   err.Error(),
			"ready":   false,
		})
		return
	}

	err = sqlDB.Ping()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "not ready",
			"message": "Database ping failed",
			"error":   err.Error(),
			"ready":   false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ready",
		"message": "Service is ready to accept traffic",
		"ready":   true,
		"time":    time.Now().Format(time.RFC3339),
	})
}

// HealthCheckLiveness checks if the service is alive
func HealthCheckLiveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "alive",
		"message": "Service is alive",
		"time":    time.Now().Format(time.RFC3339),
	})
}

// Helper functions
func getMemStats() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m
}

func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return string(rune(bytes)) + " B"
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return string(rune(bytes/div)) + " " + "KMGTPE"[exp:exp+1] + "iB"
}
