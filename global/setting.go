package global

import (
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/logger"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/setting"
)

var (
	ServerSetting    *setting.ServerSettings
	AppSetting       *setting.AppSettings
	DatabaseSettings *setting.DatabaseSettings

	Logger *logger.Logger
)
