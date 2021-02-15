package global

import (
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/logger"
	"felix.bs.com/felix/BeStrongerInGO/Gin-BlogService/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	JWTSetting      *setting.JWTSettings

	Logger *logger.Logger
)
