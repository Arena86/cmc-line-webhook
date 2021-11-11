package main

import (
	"cmc/cmc/middlewares"
	"cmc/cmc/routers"
	"os"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func getPort() string {
	p := os.Getenv("HOST_PORT")
	if p != "" {
		return ":" + p
	}
	return ":3040"
}

func main() {
	port := getPort()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.Use(location.Default())
	r.Use(middlewares.CORSMiddleware())
	rg := r.Group("cmc/v1")
	rg.Use(middlewares.CORSMiddleware())
	{
		routers.LineAPIRoute(rg)
	}
	r.Run(port)
}
