package rest

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

type SettingRestApi struct {
	Environment string `json:"environment"`
	Host        string `json:"host"`
	Port        string `json:"port"`
}

func init() {}

func Load(setting SettingRestApi, routers ...func(group *gin.RouterGroup)) {
	listenAddress := fmt.Sprintf(":%s", setting.Port)
	router := gin.New()

	// setting router
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		AllowHeaders:     []string{"Origin", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	}))
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	//router.MaxMultipartMemory = int64(constant.MAX_FILE_SIZE) << 20

	// set public folder
	//router.Static("/assets", constant.PUBLIC_ASSETS)

	apiRouters := router.Group(os.Getenv("API_VERSION"))

	//load router
	if len(routers) > 0 {
		for _, r := range routers {
			r(apiRouters)
		}
	}

	//run
	router.Run(listenAddress)
}
