package main

import (
	"realimage/delivery"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Real Image 2016
//	@version		1.0
//	@description	API for real image challenge

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8085
//	@BasePath		/

// @schemes	http
func main() {

	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/addcontributer", delivery.AddContributor)
	router.GET("/locations", delivery.GetLocations)
	router.POST("/setexludedlocations", delivery.SetExludedLocations)
	router.POST("/setincludedlocations", delivery.SetIncludedLocations)
	router.GET("/checkpermission", delivery.CheckPermission)

	router.Run(":8085")
}
