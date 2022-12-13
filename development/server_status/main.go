package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"path/filepath"
)

var dockerComposeYamlPath = "/game/docker-compose-dev.yaml"

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusInternalServerError, "not found error")
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "example.com",
	}))
	authorized.GET("/restart", func(c *gin.Context) {
		out, err := exec.Command("docker-compose", "-f", dockerComposeYamlPath, "restart").Output()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Out:\n\n%s\nErr:\n%v\n\n", string(out), err))
		} else {
			c.String(http.StatusOK, fmt.Sprintf("Out:\n\n%s\nErr:\n%v\n\n", string(out), err))
		}

	})
	authorized.GET("/stop", func(c *gin.Context) {
		out, err := exec.Command("docker-compose", "-f", dockerComposeYamlPath, "stop").Output()
		if err != nil {
			c.String(500, fmt.Sprintf("Out:\n\n%s\nErr:\n%v\n\n", string(out), err))
		} else {
			c.String(http.StatusOK, fmt.Sprintf("Out:\n\n%s\nErr:\n%v\n\n", string(out), err))
		}
	})
	authorized.GET("/pull", func(c *gin.Context) {
		out, err := exec.Command("git", "-C", filepath.Dir(dockerComposeYamlPath), "pull").Output()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Out:\n\n%s\nErr:\n%v\n\n", string(out), err))
		} else {
			c.String(http.StatusOK, fmt.Sprintf("Out:\n\n%s\nErr:\n%v\n\n", string(out), err))
		}
	})
	authorized.GET("/logs", func(c *gin.Context) {
		out, err := exec.Command("docker-compose", "-f", dockerComposeYamlPath, "logs", "--since", "30m").Output()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Out:\n\n%s\nErr:\n%v\n\n", string(out), err))
		} else {
			c.String(http.StatusOK, fmt.Sprintf("Out:\n\n%s\nErr:\n%v\n\n", string(out), err))
		}
	})
	authorized.GET("/status", func(c *gin.Context) {
		out, err := exec.Command("docker-compose", "-f", dockerComposeYamlPath, "ps", "-a").Output()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Out:\n\n%s\nErr:\n%v\n\n", string(out), err))
		} else {
			c.String(http.StatusOK, fmt.Sprintf("Out:\n\n%s\nErr:\n%v\n\n", string(out), err))
		}
	})
	authorized.GET("/help", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"restart": "Restart all containers",
			"stop":    "Stop all containers",
			"pull":    "Pull latest code",
			"logs":    "Get logs",
			"status":  "Get status",
			"help":    "Get help",
		})
	})
	if err := r.Run("172.17.0.1:18080"); err != nil {
		fmt.Println(err)
		return
	}
}
