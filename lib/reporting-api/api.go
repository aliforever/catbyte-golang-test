package reporting_api

import (
	"catbyte-golang-test/shared/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type reportingAPI struct {
	server *gin.Engine
	rc     *redis.Client
	logger logger.Logger
}

func NewReportingAPIService(rdb *redis.Client, l logger.Logger) *reportingAPI {
	if l == nil {
		l = logger.DefaultLogger{}
	}
	return &reportingAPI{
		server: gin.Default(),
		rc:     rdb,
		logger: l,
	}
}

func (a *reportingAPI) RegisterRouters() {
	a.server.POST("/message/list", a.ListMessagesHandler)
}

func (a *reportingAPI) Start(address ...string) error {
	a.RegisterRouters()

	return a.server.Run(address...)
}
