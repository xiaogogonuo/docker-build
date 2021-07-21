package app

import (
	"docker-build/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Run(v *viper.Viper) {
	r := gin.Default()

	r.GET("/server2", func(context *gin.Context) {
		logger.Info("server2 running")
		context.String(http.StatusOK, v.GetString("name"))
	})

	if err := r.Run(); err != nil {
		logger.Fatal(err.Error())
	}
}