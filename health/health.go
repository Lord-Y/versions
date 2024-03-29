// Package health assemble all functions required for health checks
package health

import (
	"net/http"
	"os"
	"strings"

	"github.com/Lord-Y/versions/mysql"
	"github.com/Lord-Y/versions/postgres"
	"github.com/gin-gonic/gin"
)

// Health permit to return basic health check
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"health": "OK"})
}

// Health permit to return basic health check
func Healthz(c *gin.Context) {
	db := make(map[string]interface{})
	var count int
	s := strings.TrimSpace(os.Getenv("SQL_DRIVER"))
	if s == "mysql" {
		if mysql.Ping() {
			db["mysql"] = "OK"
			count += 1
		} else {
			db["mysql"] = "NOT OK"
		}
	} else {
		if postgres.Ping() {
			db["postgresql"] = "OK"
			count += 1
		} else {
			db["postgresql"] = "NOT OK"
		}
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{"status": db})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": db})
	}
}
