package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-api/app/config"
)

func LoadRouter(config *config.ConfigAppEnv, db *gorm.DB) {
	router := gin.Default()
	port := config.Port
	host := config.Host
	prefix := config.Prefix

	trustedProxy := "trusted_proxy_ip"
	router.SetTrustedProxies([]string{trustedProxy})

	apiV1 := router.Group("/" + prefix + "/v1")
	{
		apiV1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello API!",
			})
		})
	}

	router.Run(host + ":" + port)
}
