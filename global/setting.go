package global

import (
	"github.com/baconYao/gin-blog/pkg/logger"
	"github.com/baconYao/gin-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
