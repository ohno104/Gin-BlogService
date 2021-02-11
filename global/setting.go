package global

import (
	"felix.bs.com/felix/BeStrongerInGO/02_GinBlog/pkg/logger"
	"felix.bs.com/felix/BeStrongerInGO/02_GinBlog/pkg/setting"
)

var (
	ServerSetting    *setting.ServerSettings
	AppSetting       *setting.AppSettings
	DatabaseSettings *setting.DatabaseSettings

	Logger *logger.Logger
)
