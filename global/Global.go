package global

import (
	"context"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/object"
)

const (
	//PreVersion = "0.0.0 Build20190101"
	//TestVersion = "0.0.0 Build20190101"
	Version   = "0.0.0 Build20190101"
	Type      = "Z9DataTransHolidayReport"
	SecretKey = "Z9DataTransHolidayReport"
)

var Ctx context.Context
var Cancel func()

//程序启动参数
var Args *object.ProgramArgs

//系统参数
var SysConfig *object.SystemConfig
